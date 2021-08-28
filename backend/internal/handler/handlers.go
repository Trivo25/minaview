package handler

import (
	"encoding/json"
	"github.com/trivo25/mina-view/backend/internal/db"
	"net/http"
)

func GetterCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categories, err := db.QueryCategories()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Error: err.Error(),
			Data:  nil,
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Error: "",
			Data:  categories,
		})
	}
}
func GetterServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	services, err := db.QueryServices()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Error: err.Error(),
			Data:  nil,
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Error: "",
			Data:  services,
		})
	}
}
