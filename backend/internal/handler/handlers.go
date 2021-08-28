package handler

import (
	"encoding/json"
	"github.com/trivo25/mina-view/backend/internal/db"
	"net/http"
)

func GetterCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	db.QueryServices()
	json.NewEncoder(w).Encode(Response{
		Error: "Some error has occured",
		Data:  struct{}{},
	})
}
