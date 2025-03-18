package security

type Hash interface {
	Hash(password string) (string, error)
	Compare(hashed string, actual string) error
}
