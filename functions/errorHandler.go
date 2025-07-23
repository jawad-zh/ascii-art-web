package functions

import (
	"html/template"
	"net/http"
	"strconv"
)

// Error page handler
func ErrorHandler(w http.ResponseWriter, text string, state int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError) //hada
		return
	}

	data := map[string]string{
		"errorText": text,
		"statue":    strconv.Itoa(state),
	}

	w.WriteHeader(state) //
	tmpl.Execute(w, data) //
}
