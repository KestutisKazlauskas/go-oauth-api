package access_token

import (
	"github.com/KestutisKazlauskas/go-oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type serivce struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &serivce{
		repository: repo,
	}
}

func (s *serivce) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("No access_token_id provided.")
	}
	accessToken, err := s.repository.GetById(accessTokenId)

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}