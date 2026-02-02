package main

import (
	"go-crud/internal/config"
	"go-crud/internal/controller"
	"go-crud/internal/db"
	"go-crud/internal/repository"
	"go-crud/internal/router"
	"go-crud/internal/service"
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
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)
	categoryController := controller.NewCategoryController(categoryService)
	mux := router.New(userController, productController, categoryController)

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
