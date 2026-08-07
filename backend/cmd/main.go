package main 

import (
	"tm/internal/repository/postgres"
	"tm/internal/api/router"
	"tm/internal/api/handlers"
	"tm/internal/services"
	"net/http"
	"log"
)

func main() {
	db, err := postgres.InitDB()
	if err != nil {
		log.Fatalf("failed connect to db with error %v \n", err)
	}

	repo := postgres.NewRepo(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)
	rout := router.NewRouter(handler)

	http.ListenAndServe(":8080", rout)
}