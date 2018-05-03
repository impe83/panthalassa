package panthalassa

import (
	"errors"
	"github.com/Bit-Nation/panthalassa/keyManager"
)

var panthalassaInstance *panthalassa

//Create a new panthalassa instance
func Start(accountStore, password string) error {

	//Exit if instance was already created and not stopped
	if panthalassaInstance != nil {
		return errors.New("call stop first in order to create a new panthalassa instance")
	}

	//Create key manager
	km, err := keyManager.OpenWithPassword(accountStore, password)
	if err != nil {
		return err
	}

	//Create panthalassa instance
	panthalassaInstance = &panthalassa{
		km: km,
	}

	return nil

}

//Create a new panthalassa instance with the mnemonic
func StartFromMnemonic(accountStore, mnemonic string) error {

	if panthalassaInstance != nil {
		return errors.New("call stop first in order to create a new panthalassa instance")
	}

	//Create key manager
	km, err := keyManager.OpenWithMnemonic(accountStore, mnemonic)
	if err != nil {
		return err
	}

	//Create panthalassa instance
	panthalassaInstance = &panthalassa{
		km: km,
	}

	return nil

}

//Eth Private key
func EthPrivateKey() (string, error) {

	if panthalassaInstance == nil {
		return "", errors.New("you have to start panthalassa")
	}

	return panthalassaInstance.EthereumPrivateKey()

}

func SendResponse(resp string) error {

	if panthalassaInstance == nil {
		return errors.New("you have to start panthalassa")
	}

	return nil

}

//Export the current account store with given password
func ExportAccountStore(pw, pwConfirm string) (string, error) {

	if panthalassaInstance == nil {
		return "", errors.New("you have to start panthalassa")
	}

	return panthalassaInstance.Export(pw, pwConfirm)

}

//Stop panthalassa
func Stop() error {

	//Exit if not started
	if panthalassaInstance == nil {
		return errors.New("you have to start panthalassa in order to stop it")
	}

	//Stop panthalassa
	err := panthalassaInstance.Stop()
	if err != nil {
		return err
	}

	//Reset singleton
	panthalassaInstance = nil

	return nil
}
