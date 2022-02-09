package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Encode(writer http.ResponseWriter, i interface{}) {
	err := json.NewEncoder(writer).Encode(i)
	CheckError(err)
}
