package storage

import (
)
import "github.com/lastbackend/cli/libs/io/filesystem"

type Token struct {
	Path string
}

func New(path string) *Token {
	return &Token{Path:path}
}

func (t *Token) Set(token string) (err error) {

	err = filesystem.CreateFile(t.Path)
	if err != nil {
		return err
	}
	err = filesystem.WriteStrToFile(t.Path, token)
	if err != nil {
		return err
	}
	return nil

}

func (t *Token) Delete() (err error) {

	err = t.Set("")
	if err != nil {
		return err
	}
	return nil

}

func (t *Token) Get() (token string, err error) {

	b, err := filesystem.ReadFile(t.Path)
	if err != nil {
		return "", err
	}
	token = string(b)

	return token, nil

}