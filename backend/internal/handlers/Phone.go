package handlers


import (
	"encoding/json"
	"net/http"
	"phone-accounting-system/internal/middleware"
	"phone-accounting-system/internal/models"
	"phone-accounting-system/internal/repository"
	"strconv"
)


type PhoneHandler struct {
	PhoneRepo     *repository.PhoneNumberRepo
	PhoneTypeRepo *repository.PhoneNumberTypeRepo
	Mux           *http.ServeMux
}


func (h *PhoneHandler) Init() {
	h.Mux.HandleFunc("GET /api/phones", middleware.AuthMiddleware(h.GetAllPhones))
	h.Mux.HandleFunc("GET /api/phones/{id}", middleware.AuthMiddleware(h.GetPhoneById))
	h.Mux.HandleFunc("POST /api/phones", middleware.AuthMiddleware(h.CreatePhone))
	h.Mux.HandleFunc("PUT /api/phones", middleware.AuthMiddleware(h.UpdatePhone))
	h.Mux.HandleFunc("DELETE /api/phones", middleware.AuthMiddleware(h.DeletePhone))
	h.Mux.HandleFunc("GET /api/persons/{id}/phones", middleware.AuthMiddleware(h.GetPhonesByPerson))

	h.Mux.HandleFunc("GET /api/phone-types", middleware.AuthMiddleware(h.GetAllPhoneTypes))
	h.Mux.HandleFunc("GET /api/phone-types/{id}", middleware.AuthMiddleware(h.GetPhoneTypeById))
	h.Mux.HandleFunc("POST /api/phone-types", middleware.AuthMiddleware(h.CreatePhoneType))
	h.Mux.HandleFunc("PUT /api/phone-types", middleware.AuthMiddleware(h.UpdatePhoneType))
	h.Mux.HandleFunc("DELETE /api/phone-types", middleware.AuthMiddleware(h.DeletePhoneType))
}


func (h *PhoneHandler) GetAllPhones(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	phones := h.PhoneRepo.GetAllPhoneNumbers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phones)
}


func (h *PhoneHandler) GetPhoneById(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	phone := h.PhoneRepo.GetPhoneNumberById(id)
	if phone == nil {
		http.Error(w, "Phone not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phone)
}


func (h *PhoneHandler) CreatePhone(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	var req struct {
		PhoneNumberValue  string  `json:"phone_number_value"`
		PersonId          int64   `json:"person_id"`
		PhoneNumberTypeId int64   `json:"phone_number_type_id"`
		Comment           *string `json:"comment"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newPhone := models.PhoneNumber{
		PhoneNumberValue:  req.PhoneNumberValue,
		PersonId:          req.PersonId,
		PhoneNumberTypeId: req.PhoneNumberTypeId,
		Comment:           req.Comment,
	}

	id, err := h.PhoneRepo.CreatePhoneNumber(newPhone)
	if err != nil {
		http.Error(w, "Failed to create phone", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}


func (h *PhoneHandler) UpdatePhone(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	var req struct {
		Id                int64   `json:"id"`
		PhoneNumberValue  string  `json:"phone_number_value"`
		PersonId          int64   `json:"person_id"`
		PhoneNumberTypeId int64   `json:"phone_number_type_id"`
		Comment           *string `json:"comment"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updatePhone := models.PhoneNumber{
		Id:                req.Id,
		PhoneNumberValue:  req.PhoneNumberValue,
		PersonId:          req.PersonId,
		PhoneNumberTypeId: req.PhoneNumberTypeId,
		Comment:           req.Comment,
	}

	err := h.PhoneRepo.SetPhoneNumber(updatePhone)
	if err != nil {
		http.Error(w, "Failed to update phone", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}


func (h *PhoneHandler) DeletePhone(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	var req struct {
		Id int64 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	delPhone := models.PhoneNumber{Id: req.Id}
	err := h.PhoneRepo.RemovePhoneNumber(delPhone)
	if err != nil {
		http.Error(w, "Failed to delete phone", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}


func (h *PhoneHandler) GetPhonesByPerson(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	phones := h.PhoneRepo.GetUsersPhoneNumbers(models.User{Id: id})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phones)
}


func (h *PhoneHandler) GetAllPhoneTypes(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	types := h.PhoneTypeRepo.GetAllPhoneNumberTypes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types)
}


func (h *PhoneHandler) GetPhoneTypeById(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	phoneType := h.PhoneTypeRepo.GetPhoneNumberTypeById(id)
	if phoneType == nil {
		http.Error(w, "Phone type not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phoneType)
}


func (h *PhoneHandler) CreatePhoneType(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	var req struct {
		TypeName string `json:"type_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newType := models.PhoneNumberType{TypeName: req.TypeName}
	id, err := h.PhoneTypeRepo.CreatePhoneNumberType(newType)
	if err != nil {
		http.Error(w, "Failed to create phone type", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}


func (h *PhoneHandler) UpdatePhoneType(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	var req struct {
		Id       int64  `json:"id"`
		TypeName string `json:"type_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updateType := models.PhoneNumberType{Id: req.Id, TypeName: req.TypeName}
	err := h.PhoneTypeRepo.SetPhoneNumberType(updateType)
	if err != nil {
		http.Error(w, "Failed to update phone type", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}


func (h *PhoneHandler) DeletePhoneType(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 2 && user.RoleId != 1 {
		http.Error(w, "Forbidden: admin or operator only", http.StatusForbidden)
		return
	}

	var req struct {
		Id int64 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	delType := models.PhoneNumberType{Id: req.Id}
	err := h.PhoneTypeRepo.RemovePhoneNumberType(delType)
	if err != nil {
		http.Error(w, "Failed to delete phone type", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}