package util

import (
	"log"
)

//
// CheckError logs with a fatal if err is not nil
//
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
