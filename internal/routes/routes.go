package routes

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// SetupRoutes configura as rotas para a aplicação
func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", homeHandler).Methods("GET")
}

// homeHandler exibe a página principal (index.html)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "Não foi possível carregar a página.", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
