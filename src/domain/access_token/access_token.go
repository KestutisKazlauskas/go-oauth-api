package access_token

import (
	"github.com/KestutisKazlauskas/go-oauth-api/src/utils/errors"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId int64 `json:"user_id"`
	//Diferent for devices 
	ClientId int64 `json:"client_id"`
	Expires int64 `json:"expires"`
}

func (accessToken *AccessToken) Validate() *errors.RestErr {

	if accessToken.AccessToken == "" {
		return errors.NewBadRequestError("invalid access_token")
	}

	if accessToken.UserId <= 0 {
		return errors.NewBadRequestError("invalid user_id")
	}

	if accessToken.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client_id")
	}

	if accessToken.Expires <= 0 {
		return errors.NewBadRequestError("invalid expires")
	}

	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime*time.Hour).Unix(),
	}
}

func (accessToken *AccessToken) IsExpired() bool {
	return time.Unix(accessToken.Expires, 0).Before(time.Now().UTC())
}