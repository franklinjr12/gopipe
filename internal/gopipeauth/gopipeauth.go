package gopipeauth

import (
	"encoding/binary"
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

func ExtractUserAndKeyFromBytes(payload []byte) (dataInputAuth DataInputAuth) {
	// user and apikey will always be the first 24 bytes of a packet
	dataInputAuth.UserId = uint64(binary.LittleEndian.Uint64(payload[0:8]))
	dataInputAuth.ApiKey = string(payload[8:24])
	return dataInputAuth
}
