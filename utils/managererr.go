package utils

import (
	"log"
	"os"
)

type Err struct {
	Message    string
	Reason     string
	Supplement string
}

func Print_err(err Err) {
	if err.Message != "" {
		log.Println(err.Message)
	}
	if err.Reason != "" {
		log.Println(err.Reason)
	}
	if err.Supplement != "" {
		log.Println(err.Supplement)
	}
	os.Exit(1)
}

func Check_err(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
