package handler

import (
	"encoding/json"
	"github.com/trivo25/mina-view/backend/internal/db"
	"net/http"
)

func HandleServiceRequest(w http.ResponseWriter, r *http.Request) {
	var req ServicePOSTRequest

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Error:     err.Error(),
			ErrorCode: 500,
			Data:      nil,
		})
	}

	err = db.InsertService(db.Service{
		ServiceName:        req.ServiceName,
		ServiceDescription: req.ServiceDescription,
		ServiceLogo:        req.ServiceLogo,
		ServiceWebsite:     req.ServiceWebsite,
		ServiceCreator:     req.ServiceCreator,
		ServiceContact:     req.ServiceContact,
	})

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Error:     err.Error(),
			ErrorCode: 500,
			Data:      nil,
		})
	}
	// Do something with the Person struct...

}

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
