package main

import (
	"fmt"
	"html/template"
	"net/http"

	"asciiArt/functions"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/ascii-art.html")
}

func handleAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	tmp1, err := template.ParseFiles("static/ascii-art.html")
	if err != nil {
		return
	}
	Banner := r.FormValue("select")
	asciiMap, err := functions.ReadAsciiBanner(Banner)
	// if err != nil {
	// 	t:=fmt.Sprint(err)
	// 	http.Error(w, t, http.StatusAlreadyReported)
	// 	fmt.Print("ther is no data")
	// 	return
	// }
	// if len(asciiMap) == 0 {
	// 	http.Error(w, "err", http.StatusAlreadyReported)
	// 	fmt.Print("ther is no data")
	// 	return
	// }
	result := functions.AsciiRepresentation(text, asciiMap)
	tmp1.Execute(w, map[string]string{"Result": result})
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ascii-art", handleAscii)

	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
