package security

import (
	"context"
	"rania-eskristal/src/applications/security"
	"rania-eskristal/src/commons/config"
	"rania-eskristal/src/commons/exceptions"
	"rania-eskristal/src/domains/authentications"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenManager struct {
	Config *config.JWTConfig
}

func NewJwtTokenManager(config *config.JWTConfig) security.AuthenticationTokenManager {
	return &JwtTokenManager{
		Config: config,
	}
}

func (j *JwtTokenManager) Generate(ctx context.Context, claims *authentications.AuthenticationPayload) (*string, error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(j.Config.AccessKey.ExpireTimeInHours) * time.Hour))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(j.Config.AccessKey.Key))
	if err != nil {
		return nil, err
	}

	return &ss, nil
}

func (j *JwtTokenManager) Verify(ctx context.Context, tokenString string) (*authentications.AuthenticationPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authentications.AuthenticationPayload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exceptions.NewAuthenticationError("ERR_JWT_METHOD")
		}
		return []byte(j.Config.AccessKey.Key), nil
	})

	if err != nil {
		return nil, exceptions.NewAuthenticationError("ERR_UNKWON_TOKEN")
	}

	if claims, ok := token.Claims.(*authentications.AuthenticationPayload); ok && token.Valid {
		return claims, nil

	}

	return nil, exceptions.NewAuthenticationError("ERR_PAYLOAD_TOKEN")

}
