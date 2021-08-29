package handler

import (
	"encoding/json"
	"github.com/trivo25/mina-view/backend/internal/db"
	"net/http"
)

func GetterCategories(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	categories, err := db.QueryCategories()

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Error:     err.Error(),
			ErrorCode: 500,
			Data:      nil,
		})
	} else {
		json.NewEncoder(w).Encode(Response{
			Error:     "",
			ErrorCode: 200,
			Data:      categories,
		})
	}
}
func GetterServices(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	services, err := db.QueryServices()

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Error:     err.Error(),
			ErrorCode: 500,
			Data:      nil,
		})
	} else {
		json.NewEncoder(w).Encode(Response{
			Error:     "",
			ErrorCode: 200,
			Data:      services,
		})
	}
}
