package api

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	pb "github.com/Bit-Nation/panthalassa/api/pb"
	dapp "github.com/Bit-Nation/panthalassa/dapp"
	ed25519 "golang.org/x/crypto/ed25519"
)

func (a *API) ShowModal(title, layout string, dAppIDKey ed25519.PublicKey) error {
	return a.dAppApi.ShowModal(title, layout, dAppIDKey)
}

func (a *API) SendEthereumTransaction(value, to, data string) (string, error) {
	return a.dAppApi.SendEthereumTransaction(value, to, data)
}

func (a *API) SaveDApp(dApp dapp.JsonRepresentation) error {
	return a.dAppApi.SaveDApp(dApp)
}

type DAppApi struct {
	api *API
}

// request to show a modal
func (a *DAppApi) ShowModal(title, layout string, dAppPubKey ed25519.PublicKey) error {

	// send request
	resp, err := a.api.request(&pb.Request{
		ShowModal: &pb.Request_ShowModal{
			DAppPublicKey: dAppPubKey,
			Title:         title,
			Layout:        layout,
		},
	}, time.Second*20)
	if err != nil {
		return err
	}

	// close since we don't care about the response
	resp.Closer <- nil

	return nil
}

// send an ethereum transaction to api
func (a *DAppApi) SendEthereumTransaction(value, to, data string) (string, error) {

	// send request
	resp, err := a.api.request(&pb.Request{
		SendEthereumTransaction: &pb.Request_SendEthereumTransaction{
			Value: value,
			To:    to,
			Data:  data,
		},
	}, time.Second*120)
	if err != nil {
		return "", err
	}

	ethTx := resp.Msg.SendEthereumTransaction
	if ethTx == nil {
		resp.Closer <- errors.New("got nil response")
	}

	objTx := map[string]interface{}{
		"nonce":    ethTx.Nonce,
		"gasPrice": ethTx.GasPrice,
		"gasLimit": ethTx.GasLimit,
		"to":       ethTx.To,
		"value":    ethTx.Value,
		"data":     ethTx.Data,
		"v":        ethTx.V,
		"r":        ethTx.R,
		"s":        ethTx.S,
		"chainId":  ethTx.ChainID,
		"from":     ethTx.From,
		"hash":     ethTx.Hash,
	}

	raw, err := json.Marshal(objTx)
	if err != nil {
		resp.Closer <- err
		return "", err
	}

	return string(raw), nil

}

// save DApp
func (a *DAppApi) SaveDApp(dApp dapp.JsonRepresentation) error {

	resp, err := a.api.request(&pb.Request{
		SaveDApp: &pb.Request_SaveDApp{
			AppName:          dApp.Name,
			Code:             dApp.Code,
			Signature:        hex.EncodeToString(dApp.Signature),
			SigningPublicKey: hex.EncodeToString(dApp.SignaturePublicKey),
		},
	}, time.Second*10)

	if err != nil {
		return err
	}

	resp.Closer <- nil

	return err
}
