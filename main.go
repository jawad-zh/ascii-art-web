package main

import (
	"fmt"
	"net/http"

	"asciiArt/functions"
)

func main() {
	functions.Router()
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
