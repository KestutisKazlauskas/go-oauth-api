package db

import (
	"github.com/KestutisKazlauskas/go-oauth-api/src/domain/access_token"
	"github.com/KestutisKazlauskas/go-oauth-api/src/utils/errors"
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
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {

}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError("Database connections error", err)
	}
	defer session.Close()

	accessToken := access_token.AccessToken{}
	if err = session.Query(QueryGetAccessToken, id).Scan(
		&accessToken.AccessToken, 
		&accessToken.UserId, 
		&accessToken.ClientId, 
		&accessToken.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("Record not found.")
		}
		return nil, errors.NewInternalServerError("Error on querying Casandra", err)
	}

	return &accessToken, nil
}

func (r *dbRepository) Create(accessToken access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError("Database connections error", err)
	}
	defer session.Close()

	if err = session.Query(
		QueryCreateAccessToken,
		accessToken.AccessToken,
		accessToken.UserId,
		accessToken.ClientId,
		accessToken.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError("Insert query error", err)
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(accessToken access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError("Database connections error", err)
	}
	defer session.Close()

	if err = session.Query(
		QueryUpdateExpirationTime,
		accessToken.Expires,
		accessToken.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError("Update query error", err)
	}

	return nil
}
