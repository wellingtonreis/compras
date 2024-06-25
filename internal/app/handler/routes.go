package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/home", HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/upload", UploadHandler).Methods(http.MethodPost)
	router.HandleFunc("/quotation-history", ListQuotationHistoryHandler).Methods(http.MethodPost)
	return router
}
