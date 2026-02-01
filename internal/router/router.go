package router

import (
	"go-crud/internal/controller"
	"net/http"
)

func New(userController controller.UserController) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Register specific routes first
	mux.HandleFunc("/users", userController.GetAllUsers)
	mux.HandleFunc("/users/new", userController.NewUserForm)
	mux.HandleFunc("/users/create", userController.CreateUser)
	mux.HandleFunc("/users/edit", userController.EditForm)
	mux.HandleFunc("/users/update", userController.UpdateUser)
	mux.HandleFunc("/users/delete", userController.DeleteUser)

	// Root path shows home page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			userController.Home(w, r)
			return
		}
		http.NotFound(w, r)
	})

	return mux
}
