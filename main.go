package main

import (
	"html/template"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/predator272/app-go/internal/helper"
	"github.com/predator272/app-go/internal/vacancy"
)

const (
	RouteRoot         = "/"
	RouteApiVacancies = "/api/vacancies/"
)

func Render(w http.ResponseWriter, code int, data any, name string, files ...string) error {
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}
	w.WriteHeader(code)
	return tmpl.ExecuteTemplate(w, name, data)
}

func main() {
	db := helper.ErrorArg(sqlx.Open("sqlite3", "database.db"))
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc(RouteRoot, func(w http.ResponseWriter, r *http.Request) {
		if Render(w, http.StatusOK, nil, "index.html", "resources/layout.html", "resources/index.html") != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	})

	vacancy.Register(RouteApiVacancies, mux, db)

	http.ListenAndServe("0.0.0.0:80", mux)
}
