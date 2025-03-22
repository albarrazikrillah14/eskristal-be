package security

import (
	"context"
	"rania-eskristal/src/domains/authentications"
)

type AuthenticationTokenManager interface {
	Generate(ctx context.Context, claims *authentications.AuthenticationPayload) (*string, error)
	Verify(ctx context.Context, token string) (*authentications.AuthenticationPayload, error)
}
