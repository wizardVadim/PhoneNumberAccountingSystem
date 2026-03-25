package handlers


import (
	"encoding/json"
	"log"
	"net/http"
	"phone-accounting-system/internal/middleware"
	"phone-accounting-system/internal/models"
	"phone-accounting-system/internal/repository"
	"strconv"
)


type UserHandler struct {
	UserRepo     *repository.UserRepo
	UserRoleRepo *repository.UserRoleRepo
	Mux          *http.ServeMux
}


func (h *UserHandler) Init() {
	h.Mux.HandleFunc("GET /api/users", middleware.AuthMiddleware(h.GetUsers))
	h.Mux.HandleFunc("GET /api/users/{id}", middleware.AuthMiddleware(h.GetUserById))
	h.Mux.HandleFunc("POST /api/users", middleware.AuthMiddleware(h.CreateUser))
	h.Mux.HandleFunc("PUT /api/users", middleware.AuthMiddleware(h.UpdateUser))
	h.Mux.HandleFunc("DELETE /api/users", middleware.AuthMiddleware(h.DeleteUser))

	h.Mux.HandleFunc("GET /api/roles", middleware.AuthMiddleware(h.GetAllRoles))
	h.Mux.HandleFunc("GET /api/roles/{id}", middleware.AuthMiddleware(h.GetRoleById))
	h.Mux.HandleFunc("POST /api/roles", middleware.AuthMiddleware(h.CreateRole))
	h.Mux.HandleFunc("PUT /api/roles", middleware.AuthMiddleware(h.UpdateRole))
	h.Mux.HandleFunc("DELETE /api/roles", middleware.AuthMiddleware(h.DeleteRole))
}


func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 0 {
		http.Error(w, "Forbidden: admin only", http.StatusForbidden)
		return
	}

	users := h.UserRepo.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 0 {
		http.Error(w, "Forbidden: admin only", http.StatusForbidden)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	curUser := h.UserRepo.GetUserById(id)
	if curUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(curUser)
}


func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 0 {
		http.Error(w, "Forbidden: admin only", http.StatusForbidden)
		return
	}

	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		RoleId   int64  `json:"role_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newUser := models.User{
		Login:    req.Login,
		Password: req.Password,
		RoleId:   req.RoleId,
	}

	id, err := h.UserRepo.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 0 {
		http.Error(w, "Forbidden: admin only", http.StatusForbidden)
		return
	}

	var req struct {
		Id       int64  `json:"id"`
		Login    string `json:"login"`
		Password string `json:"password"`
		RoleId   int64  `json:"role_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updateUser := models.User{
		Id:       req.Id,
		Login:    req.Login,
		Password: req.Password,
		RoleId:   req.RoleId,
	}

	err := h.UserRepo.SetUser(updateUser)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}


func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 0 {
		http.Error(w, "Forbidden: admin only", http.StatusForbidden)
		return
	}

	var req struct {
		Id int64 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	remUser := models.User{Id: req.Id}
	err := h.UserRepo.RemoveUser(remUser)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to remove user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}


func (h *UserHandler) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 0 {
		http.Error(w, "Forbidden: admin only", http.StatusForbidden)
		return
	}

	roles := h.UserRoleRepo.GetAllRoles()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}


func (h *UserHandler) GetRoleById(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 0 {
		http.Error(w, "Forbidden: admin only", http.StatusForbidden)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	role := h.UserRoleRepo.GetRoleById(id)
	if role == nil {
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}


func (h *UserHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 0 {
		http.Error(w, "Forbidden: admin only", http.StatusForbidden)
		return
	}

	var req struct {
		RoleName string `json:"role_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newRole := models.UserRole{RoleName: req.RoleName}
	id, err := h.UserRoleRepo.CreateUserRole(newRole)
	if err != nil {
		http.Error(w, "Failed to create role", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}


func (h *UserHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 0 {
		http.Error(w, "Forbidden: admin only", http.StatusForbidden)
		return
	}

	var req struct {
		Id       int64  `json:"id"`
		RoleName string `json:"role_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updateRole := models.UserRole{Id: req.Id, RoleName: req.RoleName}
	err := h.UserRoleRepo.SetUserRole(updateRole)
	if err != nil {
		http.Error(w, "Failed to update role", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}


func (h *UserHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 0 {
		http.Error(w, "Forbidden: admin only", http.StatusForbidden)
		return
	}

	var req struct {
		Id int64 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	delRole := models.UserRole{Id: req.Id}
	err := h.UserRoleRepo.RemoveUserRole(delRole)
	if err != nil {
		http.Error(w, "Failed to delete role", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}