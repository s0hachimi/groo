package groupie

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/errors.html"))

func ServeErrorPage(w http.ResponseWriter, r *http.Request, code int) {
	var message string

	switch code {
	case 400:
		message = "Bad Request"
	case 403:
		message = "Forbidden ! You don't have permission to access this resource."
	case 404:
		message = "page not found."
	case 405:
		message = "The method is not allowed for this resource."
	case 500:
		message = "There was an internal server error."
	default:
		message = "An unexpected error occurred."
	}

	w.WriteHeader(code) // Set the HTTP status code

	tmpl.Execute(w, map[string]interface{}{
		"Code":    code,
		"Message": message,
	})
}
