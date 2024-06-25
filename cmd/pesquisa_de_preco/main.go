package main

import (
	"log"
	"net/http"

	"github.com/wellingtonreis/compras/internal/app/middleware"

	"github.com/wellingtonreis/compras/configs"

	"github.com/wellingtonreis/compras/internal/app/handler"
)

func main() {
	router := handler.SetupRoutes()
	corsHandler := middleware.EnableCors(router)

	configs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	log.Println("Servidor iniciado na porta 3000...")
	if err := http.ListenAndServe(":"+configs.WebServerPort, corsHandler); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
