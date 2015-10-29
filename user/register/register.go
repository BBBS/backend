package register

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/pdxjohnny/numapp/api"
	"github.com/pdxjohnny/numapp/variables"
)

// Register takes a doc and attempts to create a new user
func Register(registerDoc map[string]interface{}) error {
	id, ok := registerDoc["username"].(string)
	if ok != true {
		return errors.New("Need a username to register")
	}
	password, ok := registerDoc["password"].([]byte)
	if ok != true {
		return errors.New("Need a password to register")
	}
	doc, err := api.GetAccount(variables.ServiceDBURL, id)
	if err != nil || doc != nil {
		return errors.New("Username is already taken")
	}

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		variables.BcryptCost,
	)
	if err != nil {
		return err
	}
	registerDoc["password"] = string(hashedPassword)

	_, err = api.SaveAccount(variables.ServiceDBURL, id, registerDoc)
	if err != nil {
		return err
	}
	return nil
}
