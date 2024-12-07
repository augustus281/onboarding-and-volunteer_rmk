package tokprovider

import (
	"github.com/cesc1802/share-module/common"
	"github.com/golang-jwt/jwt/v5"
)

type AppToken struct {
	Token          string `json:"token"`
	ExpiryInSecond int64  `json:"expiry"`
}

type AppPayload struct {
	UserID         string `json:"user_id"`
	RefreshTokenID string `json:"refresh_token_id"`
	jwt.RegisteredClaims
}

func (pl AppPayload) GetUserID() string {
	return pl.UserID
}

type TokenExtractor interface {
	Extract(token string) (common.Requester, error)
}

type TokenGenerator interface {
	Generate(payload *AppPayload, expiry int64) (*AppToken, error)
}
type TokenProvider interface {
	TokenExtractor
	TokenGenerator
}
