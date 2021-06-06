package helper

import "log"

func PanicIfError(reason string, err error) {
	if err != nil {
		log.Panicf("%v error: %v", reason, err)
	}
}
