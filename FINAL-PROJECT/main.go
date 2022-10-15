package main

import (
	"net/http"

	authcontroller ""
)

func main() {
	http.HandleFunc("/", authcontroller.index)
}
