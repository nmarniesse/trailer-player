package main

import (
	"log"
)

func logInfo(message string) {
	log.Println(message)
}

func logError(err error) {
	log.Fatal(err)
}
