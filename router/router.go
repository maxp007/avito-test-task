package router

import (
	"github.com/gorilla/mux"
	"github.com/maxp007/avito-test-task/handlers"
	"net/http"
)

func init() {
}

func GetRouter() (router *mux.Router) {

	router = mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", handlers.GetAdvertListHandler).Methods(http.MethodGet)
	api.HandleFunc("/advert/{id}/", handlers.GetAdvertHandler).Methods(http.MethodGet)

	api.HandleFunc("/create", handlers.CreateAdvertHandler).Methods(http.MethodPost)

	return
}
