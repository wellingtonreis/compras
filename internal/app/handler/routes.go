package handler

import (
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/home", HomeHandler)
	router.Post("/upload", UploadHandler)
	router.Post("/quotation-history", ListQuotationHistoryHandler)
	router.Get("/classification-segment", ListOptionsCategoryHandler)
	router.Get("/purchase-items/{quotation}", ListPurchaseItemsHandler)
	router.Put("/purchase-items/{quotation}", UpdatePurchaseItemsHandler)
	return router
}
