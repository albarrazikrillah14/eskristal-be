package security

import (
	"rania-eskristal/src/applications/security"
	"rania-eskristal/src/commons/exceptions"

	"golang.org/x/crypto/bcrypt"
)

type bcryptHash struct {
}

func NewBcryptHash() security.Hash {
	return &bcryptHash{}
}

func (b *bcryptHash) Hash(plain string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plain), 10)

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (b *bcryptHash) Compare(hashed string, actual string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(actual))

	if err != nil {
		return exceptions.NewInvariantError("ERR_PASSWORD_NOT_MATCH")
	}

	return nil
}
