package handlers

import (
	"encoding/json"
	"net/http"
	"phone-accounting-system/internal/middleware"
	"phone-accounting-system/internal/models"
	"phone-accounting-system/internal/repository"
	"strconv"
)

type PersonHandler struct {
	PersonRepo *repository.PhysicalPersonRepo
	Mux        *http.ServeMux
}

func (h *PersonHandler) Init() {
	h.Mux.HandleFunc("GET /api/persons", middleware.AuthMiddleware(h.GetAllPersons))
	h.Mux.HandleFunc("GET /api/persons/sorted", middleware.AuthMiddleware(h.GetAllPersonsSorted))
	h.Mux.HandleFunc("GET /api/persons/{id}", middleware.AuthMiddleware(h.GetPersonById))
	h.Mux.HandleFunc("POST /api/persons", middleware.AuthMiddleware(h.CreatePerson))
	h.Mux.HandleFunc("PUT /api/persons", middleware.AuthMiddleware(h.UpdatePerson))
	h.Mux.HandleFunc("DELETE /api/persons", middleware.AuthMiddleware(h.DeletePerson))
	h.Mux.HandleFunc("GET /api/persons/phone-stats", middleware.AuthMiddleware(h.GetPersonsPhoneStats))
}

func (h *PersonHandler) GetAllPersons(w http.ResponseWriter, r *http.Request) {

	persons := h.PersonRepo.GetAllPhysicalPersons()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func (h *PersonHandler) GetAllPersonsSorted(w http.ResponseWriter, r *http.Request) {

	persons := h.PersonRepo.GetAllPhysicalPersonsSortedName()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func (h *PersonHandler) GetPersonById(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	person := h.PersonRepo.GetPhysicalPersonById(id)
	if person == nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func (h *PersonHandler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 1 && user.RoleId != 3 {
		http.Error(w, "Forbidden: admin or MFC worker only", http.StatusForbidden)
		return
	}

	var req struct {
		City       string  `json:"city"`
		Address    *string `json:"address"`
		FirstName  string  `json:"first_name"`
		LastName   string  `json:"last_name"`
		SecondName *string `json:"second_name"`
		BornYear   *int16  `json:"born_year"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	newPerson := models.PhysicalPerson{
		City:       req.City,
		Address:    req.Address,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		SecondName: req.SecondName,
		BornYear:   req.BornYear,
	}

	id, err := h.PersonRepo.CreatePhysicalPerson(newPerson)
	if err != nil {
		http.Error(w, "Failed to create person", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func (h *PersonHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 1 && user.RoleId != 3 {
		http.Error(w, "Forbidden: admin or MFC worker only", http.StatusForbidden)
		return
	}

	var req struct {
		Id         int64   `json:"id"`
		City       string  `json:"city"`
		Address    *string `json:"address"`
		FirstName  string  `json:"first_name"`
		LastName   string  `json:"last_name"`
		SecondName *string `json:"second_name"`
		BornYear   *int16  `json:"born_year"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updatePerson := models.PhysicalPerson{
		Id:         req.Id,
		City:       req.City,
		Address:    req.Address,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		SecondName: req.SecondName,
		BornYear:   req.BornYear,
	}

	err := h.PersonRepo.SetPhysicalPerson(updatePerson)
	if err != nil {
		http.Error(w, "Failed to update person", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func (h *PersonHandler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(*models.User)
	if user.RoleId != 1 && user.RoleId != 3 {
		http.Error(w, "Forbidden: admin or MFC worker only", http.StatusForbidden)
		return
	}

	var req struct {
		Id int64 `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	delPerson := models.PhysicalPerson{Id: req.Id}
	err := h.PersonRepo.RemovePhysicalPerson(delPerson)
	if err != nil {
		http.Error(w, "Failed to delete person", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func (h *PersonHandler) GetPersonsPhoneStats(w http.ResponseWriter, r *http.Request) {

	stats := h.PersonRepo.GetPhysicalPersonsPhoneNumbersQuantity()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
