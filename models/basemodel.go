package models

import (
	"houseduino-be/lang"
	"os"
)

// ErrHandler returns error message bassed on env debug
func ErrHandler(err error) string {
	var errmessage string
	if os.Getenv("DEBUG") == "true" {
		errmessage = err.Error()
	} else {
		errmessage = lang.Get("something_went_wrong")
	}
	return errmessage
}
