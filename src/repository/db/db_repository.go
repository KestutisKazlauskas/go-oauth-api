package db

import (
	"github.com/KestutisKazlauskas/go-oauth-api/src/domain/access_token"
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
	"github.com/KestutisKazlauskas/go-utils/logger"
	"github.com/KestutisKazlauskas/go-oauth-api/src/clients/cassandra"
	"github.com/gocql/gocql"
)

const (
	QueryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	QueryCreateAccessToken = "INSERT INTO access_tokens (access_token, client_id, user_id, expires) VALUES (?, ?, ?, ?);"
	QueryUpdateExpirationTime = "UPDATE access_tokens SET expires=? WHERE access_token=?";
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	Create(access_token.AccessToken) *rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *rest_errors.RestErr
}

type dbRepository struct {

}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *rest_errors.RestErr) {
	accessToken := access_token.AccessToken{}
	if err := cassandra.GetSession().Query(QueryGetAccessToken, id).Scan(
		&accessToken.AccessToken, 
		&accessToken.UserId, 
		&accessToken.ClientId, 
		&accessToken.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("Record not found.")
		}
		return nil, rest_errors.NewInternalServerError("Error on querying Casandra", err, logger.Log)
	}

	return &accessToken, nil
}

func (r *dbRepository) Create(accessToken access_token.AccessToken) *rest_errors.RestErr {
	if err := cassandra.GetSession().Query(
		QueryCreateAccessToken,
		accessToken.AccessToken,
		accessToken.UserId,
		accessToken.ClientId,
		accessToken.Expires,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError("Insert query error", err, logger.Log)
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(accessToken access_token.AccessToken) *rest_errors.RestErr {
	if err := cassandra.GetSession().Query(
		QueryUpdateExpirationTime,
		accessToken.Expires,
		accessToken.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError("Update query error", err, logger.Log)
	}

	return nil
}
