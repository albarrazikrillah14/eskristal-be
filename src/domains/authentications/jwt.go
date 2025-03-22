package authentications

import (
	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationPayload struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
	jwt.RegisteredClaims
}
