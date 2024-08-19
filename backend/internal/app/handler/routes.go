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
	router.Put("/quotation-history/classification-segment/{quotation}/put", UpdateClassificationSegment)
	router.Get("/classification-segment", ListOptionsCategoryHandler)
	router.Get("/purchase-items/{quotation}/get", ListPurchaseItemsHandler)
	router.Put("/purchase-items/{quotation}/put", UpdatePurchaseItemsJustifyHandler)
	router.Put("/purchase-items/{quotation}/delete", DeletePurchaseItemsHandler)
	return router
}
