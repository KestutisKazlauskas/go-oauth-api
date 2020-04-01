package rest

import (
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
	"github.com/KestutisKazlauskas/go-utils/logger"
	"github.com/KestutisKazlauskas/go-oauth-api/src/domain/users"
	"github.com/federicoleon/golang-restclient/rest"
	"time"
	"encoding/json"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

func NewRepository() UsersRepository {
	return &usersRepository{}
}

type UsersRepository interface {
	Login(string, string) (*users.User, *rest_errors.RestErr)
}

type usersRepository struct {

}

func (r *usersRepository) Login(email string, password string) (*users.User, *rest_errors.RestErr) {
	request := users.LoginRequest{
		Email: email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	
	//Tiemout happens
	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("Timeout on users api.", nil, logger.Log)
	}

	//Some errors hapens
	if response.StatusCode > 299 {
		var restErr rest_errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("Cant parse the error.", nil, logger.Log)
		}

		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("Cant parse the user response.", nil, logger.Log)
	}

	return &user, nil 
}