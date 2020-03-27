package access_token

import (
	"github.com/KestutisKazlauskas/go-oauth-api/src/utils/errors"
	"github.com/KestutisKazlauskas/go-oauth-api/src/repository/rest"
	"github.com/KestutisKazlauskas/go-oauth-api/src/repository/db"
	"github.com/KestutisKazlauskas/go-oauth-api/src/domain/access_token"
	"strings"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken,*errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type serivce struct {
	restUserRepo rest.UsersRepository
	dbRepo db.DbRepository
}

func NewService(userRepo rest.UsersRepository, dbRepo db.DbRepository) Service {
	return &serivce{
		restUserRepo: userRepo,
		dbRepo: dbRepo,
	}
}

func (s *serivce) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("No access_token_id provided.")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *serivce) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	//TODO add AuthType for credentials (Grandtype on OAuth)
	//https://www.oauth.com/oauth2-servers
	//authenticate user using User api
	user, err := s.restUserRepo.Login(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	accessToken := access_token.GetNewAccessToken(user.Id)
	accessToken.Generate()

	if err := s.dbRepo.Create(accessToken); err != nil {
		return nil, err
	}

	return &accessToken, nil
}

func (s *serivce) UpdateExpirationTime(accessToken access_token.AccessToken) *errors.RestErr {
	if err := accessToken.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(accessToken)
}