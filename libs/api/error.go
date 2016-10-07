package api

import (
	"errors"
)

type failureS struct {
	Error struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	} `json:"error"`
}

func checkError(failure *failureS) (ok bool, err error) {

	if failure.Error.Code == "" && failure.Error.Description == ""{
		return false, nil
	}
	return true, errors.New(failure.Error.Code)

}
