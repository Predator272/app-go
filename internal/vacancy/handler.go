package vacancy

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/predator272/app-go/internal/helper"
)

func RenderModel(w http.ResponseWriter, model *Model, errorCode int) {
	if model == nil {
		http.Error(w, "", errorCode)
		return
	}
	json.NewEncoder(w).Encode(&model)
}

func Register(prefix string, m *http.ServeMux, db *sqlx.DB) {
	storage := Storage{db: db}
	storage.Create()

	router := mux.NewRouter()
	router.StrictSlash(true)
	subrouter := router.PathPrefix(prefix).Subrouter()

	subrouter.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(storage.All())
	}).Methods(http.MethodGet)

	subrouter.HandleFunc("/show/{ID:\\d+}", func(w http.ResponseWriter, r *http.Request) {
		load := Model{}
		if helper.ParseUint(&load.ID, mux.Vars(r)["ID"]) != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		RenderModel(w, storage.FindById(load), http.StatusNotFound)
	}).Methods(http.MethodGet)

	subrouter.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		load := Model{}
		if json.NewDecoder(r.Body).Decode(&load) != nil || load.Validate() == false {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		RenderModel(w, storage.Save(load), http.StatusInternalServerError)
	}).Methods(http.MethodPost)

	subrouter.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		load := Model{}
		if json.NewDecoder(r.Body).Decode(&load) != nil || load.Validate() == false {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		RenderModel(w, storage.UpdateById(load), http.StatusNotFound)
	}).Methods(http.MethodPut)

	subrouter.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		load := Model{}
		if json.NewDecoder(r.Body).Decode(&load) != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		RenderModel(w, storage.DeleteById(load), http.StatusNotFound)
	}).Methods(http.MethodDelete)

	m.Handle(prefix, router)
}
