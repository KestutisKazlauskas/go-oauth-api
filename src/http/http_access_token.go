package http

import (
	"github.com/KestutisKazlauskas/go-oauth-api/src/domain/access_token"
	"github.com/KestutisKazlauskas/go-oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func  (handler *accessTokenHandler) GetById(context *gin.Context) {

	accessToken, err := handler.service.GetById(context.Param("access_token_id"))
	if err != nil {
		context.JSON(err.Status, err)
		return
	}

	context.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(context *gin.Context) {
	var accessToken access_token.AccessToken

	if err := context.ShouldBindJSON(&accessToken); err != nil {
		restErr := errors.NewBadRequestError("Invalid data provided.")
		context.JSON(restErr.Status, restErr)
		return
	} 

	if err := handler.service.Create(accessToken); err != nil {
		context.JSON(err.Status, err)
		return
	}
	context.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func (handler *accessTokenHandler) UpdateExpirationTime(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]string{"message": "success"})
}