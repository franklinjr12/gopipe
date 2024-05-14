package gopipeauth

import (
	"errors"
	"gopipe/internal/database"
)

const UNAUTHORIZED_ERROR = "Unauthorized"

type DataInputAuth struct {
	UserId uint64
	ApiKey string
}

func AuthenticateDataInput(authData DataInputAuth) error {
	if authData.UserId == 0 {
		return errors.New(UNAUTHORIZED_ERROR)
	}
	if authData.ApiKey == "" {
		return errors.New(UNAUTHORIZED_ERROR)
	}
	db := database.Open()
	if database.UserExist(db, authData.UserId, authData.ApiKey) {
		return nil
	}
	return errors.New(UNAUTHORIZED_ERROR)
}
