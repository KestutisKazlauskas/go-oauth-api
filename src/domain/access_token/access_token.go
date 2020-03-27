package access_token

import (
	"github.com/KestutisKazlauskas/go-oauth-api/src/utils/errors"
	"github.com/KestutisKazlauskas/go-oauth-api/src/utils/crypto_utils"
	"time"
	"fmt"
)

const (
	expirationTime = 24
	authTypePassword = "password"
	authTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	AuthType string `json:"auth_type"`
	Scope string `json:"scope"`

	//AuthType "password"
	Email string `json:"email"`
	Password string `json:"password"`

	//AuthType "clien_credentials"
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (accessTokenReq *AccessTokenRequest) Validate() *errors.RestErr {
	switch accessTokenReq.AuthType {
	case authTypePassword:
		break
	case authTypeClientCredentials:
		return errors.NewBadRequestError("AuthType not implemented yet")
	default:
		return errors.NewBadRequestError("invalid auth type")
	}
	return nil 
}

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

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId: userId,
		Expires: time.Now().UTC().Add(expirationTime*time.Hour).Unix(),
	}
}

func (accessToken *AccessToken) IsExpired() bool {
	return time.Unix(accessToken.Expires, 0).Before(time.Now().UTC())
}

func (accessToken *AccessToken) Generate() {
	accessToken.AccessToken = crypto_utils.GetMd5Hash(fmt.Sprintf("at-%d-%d-ran", accessToken.UserId, accessToken.Expires))
}