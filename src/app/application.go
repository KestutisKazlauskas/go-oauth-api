package app

import (
	"github.com/KestutisKazlauskas/go-oauth-api/src/domain/access_token"
	"github.com/KestutisKazlauskas/go-oauth-api/src/repository/db"
	"github.com/KestutisKazlauskas/go-oauth-api/src/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	//This is a controller
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.Create)

	router.Run(":8080")
}