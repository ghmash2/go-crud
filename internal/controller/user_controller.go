package controller

import (
	"go-crud/internal/models"
	"go-crud/internal/service"
	"go-crud/internal/validation"
	"html/template"
	"net/http"
	"strconv"
)

func NewUserController(userService service.UserService, tmpl *template.Template) UserController {
	return UserController{
		UserService: userService,
		Template:    tmpl,
	}
}

type UserController struct {
	UserService service.UserService
	Template    *template.Template
}

type viewData struct {
	Title   string
	Error   string
	Success string
	Any     any
}

func parseID(r *http.Request) (uint64, error) {
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c UserController) Home(w http.ResponseWriter, r *http.Request) {
	err := c.Template.ExecuteTemplate(w, "home", nil)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c UserController) NewUserForm(w http.ResponseWriter, r *http.Request) {
	err := c.Template.ExecuteTemplate(w, "users-new", models.User{})
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	user := validation.UserValidation{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if err := validation.ValidateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := c.UserService.CreateUser(r.Context(), user)
	if err != nil {
		err = c.Template.ExecuteTemplate(w, "users-new", models.User{
			Name:  user.Name,
			Email: user.Email,
		})
		if err != nil {
			http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (c UserController) EditForm(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := c.UserService.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = c.Template.ExecuteTemplate(w, "users-edit", user)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	id, err := parseID(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := validation.UserValidation{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if err := validation.ValidateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.UserService.UpdateUser(r.Context(), id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (c UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(r.FormValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = c.UserService.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func (c UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserService.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.Template.ExecuteTemplate(w, "users-index", users)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
