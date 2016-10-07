package errors

import "errors"

var storageErrors []map[string]string

func init()  {
	storageErrors = append(storageErrors, storageCustomErrors,
																				storageServiceErrors,
																				storageAuthErrors)
}

func Err(err error) string {

	for _, elementErrors := range storageErrors {
		internal_err, ok := elementErrors[err.Error()]
		if ok {
			return internal_err
		}
	}

	return "Unknown error"

}

func New(message string) error {

	return errors.New(message)

}