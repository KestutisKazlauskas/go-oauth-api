package access_token

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "timeExpiration constant not 24.")
}

func TestGetNewAccessToken(t *testing.T) {
	access_token := GetNewAccessToken()

	if access_token.IsExpired() {
		t.Error("New token could not be expired.")
	}

	if access_token.AccessToken != "" {
		t.Error("AcceessToken could not be empty.")
	}

	if access_token.UserId != 0 {
		t.Error("new access_token should not have UserId.")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	access_token := AccessToken{}

	if !access_token.IsExpired() {
		t.Error("Empty access_token should be expired by default.")
	}

	access_token.Expires = time.Now().UTC().Add(3*time.Hour).Unix()

	if access_token.IsExpired() {
		t.Error("access_token created 3 hours from now should be expired")
	}

}