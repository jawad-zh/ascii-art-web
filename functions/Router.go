package functions

import (
	"html/template"
	"net/http"
	"strconv"
)

func Router() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ascii-art", handleAscii)
}

func handleAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandlError(w,"method not allowd",http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	tmp1, err := template.ParseFiles("static/ascii-art.html")
	if err != nil {
		return
	}
	Banner := r.FormValue("select")
	asciiMap, err := ReadAsciiBanner(Banner)
	if err != nil {
		return
	}
	result := AsciiRepresentation(text, asciiMap)
	tmp1.Execute(w, map[string]string{"Result": result})
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.ServeFile(w, r, "static/error.html")
		HandlError(w,"method not allowd",http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
	HandlError(w,"page not found",http.StatusNotFound)
	}else{

		http.ServeFile(w, r, "static/ascii-art.html")
	}
}

func HandlError(w http.ResponseWriter, s string,n int) {
	temp1, err := template.ParseFiles("static/error.html")
	if err != nil {
		http.Error(w, "server error:", 500)
	}
	ana := strconv.Itoa(n)
		temp1.Execute(w, map[string]string{"Status": ana ,"Message":s})
}
