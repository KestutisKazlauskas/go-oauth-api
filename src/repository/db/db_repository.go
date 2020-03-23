package db

import (
	"github.com/KestutisKazlauskas/go-oauth-api/src/domain/access_token"
	"github.com/KestutisKazlauskas/go-oauth-api/src/utils/errors"
	"github.com/KestutisKazlauskas/go-oauth-api/src/clients/cassandra"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {

}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError("No database", err)
	}
	defer session.Close()

	return nil, errors.NewInternalServerError("No database", nil)
}
