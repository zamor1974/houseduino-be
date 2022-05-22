package models

import (
	"log"
)

// ErrHandler returns error message bassed on env debug
func ErrHandler(err error) string {
	var errmessage string
	errmessage = err.Error()
	/* 	if os.Getenv("DEBUG") == "true" {
	   		errmessage = err.Error()
	   	} else {
	   		errmessage = lang.Get("something_went_wrong")
	   	} */
	return errmessage
}
func PrintErrorLog(area string, err error) {
	log.Printf("%s -> Errore: %s", area, ErrHandler(err))

}
func PrintLog(area string, message string) {
	log.Printf("%s -> %s", area, message)

}
func PrintStringErrorLog(area string, messaggio string) {
	log.Printf("%s -> Errore: %s", area, messaggio)

}
