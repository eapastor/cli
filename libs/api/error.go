package api

import (
	"fmt"
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
	return true, errors.New(fmt.Sprintf("%s: %s", failure.Error.Code, failure.Error.Description))

}
