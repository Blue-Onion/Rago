package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-chi/cors"
)

const PORT = "9000"

func main() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	server:=http.Server{
		Handler: router,
		Addr: PORT,
	}
	fmt.Println("Hello")
	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err.Error())
	}
}
