package helpers

import "log"

func FailOnError(err error, str string) {
	if err != nil {
		log.Fatalf("Error on : %s", str)
	}
}
