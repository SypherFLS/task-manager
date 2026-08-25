package main

import (
	"log"
	"net/http"
	"tm/internal/api/handlers"
	"tm/internal/api/router"
	"tm/internal/auth"
	"tm/internal/repository/postgres"
	"tm/internal/services"
)

func main() {
	db, err := postgres.InitDB()
	if err != nil {
		log.Fatalf("failed connect to db with error %v \n", err)
	}

	JWTManager := auth.NewJWTManager(string(auth.SecretKey)) // заменить на конфиг

	repo := postgres.NewRepo(db)
	service := services.NewService(repo, JWTManager)
	handler := handlers.NewHandler(service)
	rout := router.NewRouter(handler, JWTManager)

	http.ListenAndServe(":8080", rout)
}
