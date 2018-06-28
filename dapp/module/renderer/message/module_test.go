package dapp

import (
	"testing"

	log "github.com/op/go-logging"
	otto "github.com/robertkrimen/otto"
	require "github.com/stretchr/testify/require"
)

func TestModule_RenderMessageError(t *testing.T) {

	l := log.MustGetLogger("")

	vm := otto.New()

	m := New(l)
	require.Nil(t, m.Register(vm))

	vm.Call("setMessageHandler", vm, func(payload otto.Value, cb otto.Value) otto.Value {

		cb.Call(cb, "I am an error")

		return otto.Value{}

	})

	_, err := m.RenderMessage(`{}`)
	require.EqualError(t, err, "I am an error")

}

func TestModule_RenderMessageSuccess(t *testing.T) {

	l := log.MustGetLogger("")

	vm := otto.New()

	m := New(l)
	require.Nil(t, m.Register(vm))

	vm.Call("setMessageHandler", vm, func(payload otto.Value, cb otto.Value) otto.Value {

		msg, err := payload.Object().Get("message")
		if err != nil {
			panic(err)
		}

		if !msg.IsObject() {
			panic("Expected message to be an object")
		}

		cb.Call(cb, nil, "{}")

		return otto.Value{}

	})

	layout, err := m.RenderMessage(`{message: {}, context: {}}`)
	require.Nil(t, err)
	require.Equal(t, "{}", layout)

}
