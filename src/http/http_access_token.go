package http

import (
	"github.com/KestutisKazlauskas/go-oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
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