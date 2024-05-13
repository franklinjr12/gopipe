package gopipeauth

import "errors"

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
	return nil
}
