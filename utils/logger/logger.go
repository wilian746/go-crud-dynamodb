package logger

import "log"

func PANIC(message string, err error) {
	if err != nil {
		log.Panic(message, err)
	}
}
