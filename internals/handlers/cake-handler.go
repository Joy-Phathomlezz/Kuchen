package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Joy-Phathomlezz/Kuchen/internals/models"
	"github.com/Joy-Phathomlezz/Kuchen/internals/models/repository"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CakeHandler struct {
	db *gorm.DB
}

func NewCakeHandler(db *gorm.DB) *CakeHandler {
	return &CakeHandler{db: db}
}

func (h *CakeHandler) CreateCake(w http.ResponseWriter, r *http.Request) {
	var cake models.Cake
	json.NewDecoder(r.Body).Decode(&cake)
	// Use the DB instance (h.db) to create a new cake via repository function
	err := repository.CreateCake(h.db, &cake)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cake)
}

func (h *CakeHandler) GetCakes(w http.ResponseWriter, r *http.Request) {
	// Use the DB instance (h.db) to retrieve all cakes via repository function
	cakes, err := repository.GetCakes(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cakes)
}

func (h *CakeHandler) GetCakeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["CakeId"])
	// Use the DB instance (h.db) to retrieve a cake by ID via repository function
	cake, err := repository.GetCakeByID(h.db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cake)
}

func (h *CakeHandler) UpdateCake(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["CakeId"])

	// Define a struct for input to avoid decoding into full Cake model
	var input struct {
		Size  string  `json:"size"`
		Price float64 `json:"price"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	// Prepare updates map with only provided fields
	updates := make(map[string]interface{})
	if input.Size != "" {
		updates["size"] = input.Size
	}
	if input.Price != 0 {
		updates["price"] = input.Price
	}

	// Use the DB instance (h.db) to update the cake via repository function
	err := repository.UpdateCake(h.db, uint(id), updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch the updated cake to return it
	cake, err := repository.GetCakeByID(h.db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cake)
}

func (h *CakeHandler) DeleteCake(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["CakeId"])
	// Use the DB instance (h.db) to delete the cake by ID via repository function
	err := repository.DeleteCake(h.db, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
