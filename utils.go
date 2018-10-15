package main

import (
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func checkHealth(w http.ResponseWriter, r *http.Request) {
	message := "All set!\n"

	w.Write([]byte(message))
}
