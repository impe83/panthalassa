package message

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	log "github.com/ipfs/go-log"
	ed25519 "golang.org/x/crypto/ed25519"

	reqLim "github.com/Bit-Nation/panthalassa/dapp/request_limitation"
	validator "github.com/Bit-Nation/panthalassa/dapp/validator"
	db "github.com/Bit-Nation/panthalassa/db"
	logger "github.com/op/go-logging"
	otto "github.com/robertkrimen/otto"
	uuid "github.com/satori/go.uuid"
)

type Module struct {
	chatStorage db.ChatStorage
	dAppPubKey  ed25519.PublicKey
	logger      *logger.Logger
	reqLim      *reqLim.CountThrottling
}

var sysLog = log.Logger("messsage")

func New(msgStorage db.ChatStorage, dAppPubKey ed25519.PublicKey, logger *logger.Logger) *Module {
	return &Module{
		chatStorage: msgStorage,
		dAppPubKey:  dAppPubKey,
		logger:      logger,
		reqLim:      reqLim.NewCountThrottling(4, time.Second*60, 10, errors.New("send message queue is full")),
	}
}

func hasKey(stack []string, search string) bool {
	for _, v := range stack {
		if v == search {
			return true
		}
	}
	return false
}

func (m *Module) Close() error {
	return nil
}

func (m *Module) Register(vm *otto.Otto) error {
	return vm.Set("sendMessage", func(call otto.FunctionCall) otto.Value {

		sysLog.Debug("send message")
		// validate function call
		v := validator.New()
		// first param must be the chat
		v.Set(0, &validator.TypeString)
		// second param must be the message payload
		v.Set(1, &validator.TypeObject)
		// callback
		v.Set(2, &validator.TypeFunction)
		cb := call.Argument(2)
		// utils to handle an occurred error
		handleError := func(errMsg string) otto.Value {
			if cb.IsFunction() {
				_, err := cb.Call(cb, errMsg)
				if err != nil {
					m.logger.Error(err.Error())
				}
				return otto.Value{}
			}
			m.logger.Error(errMsg)
			return otto.Value{}
		}
		if err := v.Validate(vm, call); err != nil {
			// in the case an callback has been passed we want to call it with the error
			return handleError(err.String())
		}

		messagePayload := call.Argument(1).Object()

		objv := validator.NewObjValidator()
		objv.Set("shouldSend", validator.ObjTypeBool, true)
		objv.Set("params", validator.ObjTypeObject, false)
		objv.Set("type", validator.ObjTypeString, false)
		if err := objv.Validate(vm, *messagePayload); err != nil {
			return handleError(err.String())
		}

		dAppMessage := db.DAppMessage{
			DAppPublicKey: m.dAppPubKey,
			Params:        map[string]interface{}{},
		}

		// chat in which the message should be persisted
		chatStr := call.Argument(0).String()
		chat, err := hex.DecodeString(chatStr)
		if err != nil {
			return handleError(err.Error())
		}
		if len(chat) != 32 {
			return handleError("chat must be 32 bytes long")
		}

		// should this message be sent or kept locally?
		shouldSendVal, err := messagePayload.Get("shouldSend")
		if err != nil {
			return handleError(err.Error())
		}
		dAppMessage.ShouldSend, err = shouldSendVal.ToBoolean()
		if err != nil {
			return handleError(err.Error())
		}

		// set optional type of the message
		if hasKey(messagePayload.Keys(), "type") {
			msgType, err := messagePayload.Get("type")
			if err != nil {
				return handleError(err.Error())
			}
			dAppMessage.Type = msgType.String()
		}

		// set optional params
		if hasKey(messagePayload.Keys(), "params") {

			// fetch params object
			paramsObj, err := messagePayload.Get("params")
			if err != nil {
				return handleError(err.Error())
			}
			params := paramsObj.Object()

			// transform params to map
			for _, objKey := range params.Keys() {
				vmValue, err := params.Get(objKey)
				value, err := vmValue.Export()
				if err != nil {
					return handleError(err.Error())
				}
				dAppMessage.Params[objKey] = value
			}

			// marshal params
			marshaledParams, err := json.Marshal(dAppMessage.Params)
			if err != nil {
				return handleError(err.Error())
			}

			// make sure it's less than 64 KB
			if len(marshaledParams) > 64*1024 {
				return handleError("the message params can't be bigger than 64 kb")
			}
			if err := json.Unmarshal(marshaledParams, &dAppMessage.Params); err != nil {
				return handleError(err.Error())
			}

		}

		// persist message
		m.reqLim.Exec(func(dec chan struct{}) {
			defer func() {
				dec <- struct{}{}
			}()

			chat, err := m.chatStorage.GetChatByPartner(chat)
			if err != nil {
				handleError(err.Error())
				return
			}
			if chat == nil {
				handleError("chat doesn't exist")
				return
			}

			id, err := uuid.NewV4()
			if err != nil {
				handleError(err.Error())
				return
			}

			err = chat.PersistMessage(db.Message{
				ID:        id.String(),
				Version:   1,
				Status:    db.StatusPersisted,
				Received:  false,
				DApp:      &dAppMessage,
				CreatedAt: time.Now().UnixNano(),
			})
			if err != nil {
				handleError(err.Error())
				return
			}
			if _, err := cb.Call(cb); err != nil {
				m.logger.Error(err.Error())
			}
		})

		return otto.Value{}

	})
}
