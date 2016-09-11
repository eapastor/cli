package storage

type IToken interface {
	Set(token string) (err error)
	Delete() (err error)
	Get() (token string, err error)
}
