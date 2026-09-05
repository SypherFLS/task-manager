package main

import (
	"log"
	"net/http"
	"os"
	"tm/internal/api/handlers"
	"tm/internal/api/router"
	"tm/internal/auth"
	"tm/internal/config"
	"tm/internal/repository/postgres"
	"tm/internal/services"
)

func main() {
	path := os.Getenv("CONFIG_PATH")
	cfg, erro := config.ConfigInit(path)
	if erro != nil {
		log.Fatalf("failed config init with error %v\n", erro)
	}

	db, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed connect to db with error %v \n", err)
	}


	jwtsecret := os.Getenv("JWT_SECRET")
	JWTManager := auth.NewJWTManager(jwtsecret) 

	repo := postgres.NewRepo(db)
	service := services.NewService(repo, JWTManager)
	handler := handlers.NewHandler(service)
	rout := router.NewRouter(handler, JWTManager, cfg)

	http.ListenAndServe(cfg.Server.Host, rout)
}
