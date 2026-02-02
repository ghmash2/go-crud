package router

import (
	"go-crud/internal/controller"
	"go-crud/internal/middleware"
	"net/http"
)

func New(userController controller.UserController, productController controller.ProductController, categoryController controller.CategoryController) http.Handler {
	mux := http.NewServeMux()

	// API endpoints for users
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userController.GetAllUsers(w, r)
		} else if r.Method == http.MethodPost {
			userController.CreateUser(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			userController.EditForm(w, r) // Returns single user as JSON
		} else if r.Method == http.MethodPut || r.Method == http.MethodPost {
			userController.UpdateUser(w, r)
		} else if r.Method == http.MethodDelete {
			userController.DeleteUser(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// API endpoints for products
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			productController.GetAllProducts(w, r)
		} else if r.Method == http.MethodPost {
			productController.CreateProduct(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			productController.EditForm(w, r) // Returns single product as JSON
		} else if r.Method == http.MethodPut || r.Method == http.MethodPost {
			productController.UpdateProduct(w, r)
		} else if r.Method == http.MethodDelete {
			productController.DeleteProduct(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// API endpoints for categories
	mux.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			categoryController.GetAllCategories(w, r)
		} else if r.Method == http.MethodPost {
			categoryController.CreateCategory(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			categoryController.EditForm(w, r) // Returns single category as JSON
		} else if r.Method == http.MethodPut || r.Method == http.MethodPost {
			categoryController.UpdateCategory(w, r)
		} else if r.Method == http.MethodDelete {
			categoryController.DeleteCategory(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Wrap with CORS middleware
	return middleware.CORS(mux)
}
