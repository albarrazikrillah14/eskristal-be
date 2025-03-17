package security

type Hash interface {
	Hash(password string) string
}
