package appjwt

import (
	"time"

	"github.com/cesc1802/share-module/common"
	"github.com/cesc1802/share-module/tokprovider"
	"github.com/golang-jwt/jwt/v5"
)

type jwtProvider struct {
	Secret string
}

func NewJwtProvider(secret string) *jwtProvider {
	return &jwtProvider{
		Secret: secret,
	}
}

func (p *jwtProvider) Extract(token string) (common.Requester, error) {
	result, err := jwt.ParseWithClaims(token, &tokprovider.AppPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(p.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !result.Valid {
		return nil, err
	}

	claims, ok := result.Claims.(*tokprovider.AppPayload)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func (p *jwtProvider) Generate(payload *tokprovider.AppPayload, expiryInSecond int64) (*tokprovider.AppToken, error) {
	payload.IssuedAt = jwt.NewNumericDate(time.Now().Local())
	payload.ExpiresAt = jwt.NewNumericDate(time.Now().Local().Add(time.Second * time.Duration(expiryInSecond)))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token, err := t.SignedString([]byte(p.Secret))
	if err != nil {
		return nil, err
	}

	return &tokprovider.AppToken{
		Token:          token,
		ExpiryInSecond: expiryInSecond,
	}, nil
}
