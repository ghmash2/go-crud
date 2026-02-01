package main

import (
	"go-crud/internal/config"
	"go-crud/internal/controller"
	"go-crud/internal/db"
	"go-crud/internal/repository"
	"go-crud/internal/router"
	"go-crud/internal/service"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {

	cfg := config.LoadConfig()

	db, err := db.OpenConnection(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	tmpl := template.Must(template.ParseGlob("web/templates/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("web/templates/users/*.html"))
	if tmpl == nil {
		log.Fatal("Failed to parse templates")
	}

	userController := controller.NewUserController(userService, tmpl)
	mux := router.New(userController)

	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on http://localhost:" + cfg.AppPort)
}
