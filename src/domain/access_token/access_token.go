package access_token

import (
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId int64 `json:"user_id"`
	//Diferent for devices 
	ClinetId int64 `json:"client_id"`
	Expires int64 `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime*time.Hour).Unix(),
	}
}

func (accessToken AccessToken) IsExpired() bool {
	return time.Unix(accessToken.Expires, 0).Before(time.Now().UTC())
}