package handler

import (
	"github.com/trivo25/mina-view/backend/internal/db"
)

// defines the structure of the response with some meta data and the actual data
type Response struct {
	Error     string
	ErrorCode int
	Data      interface{}
}

type ServicePOSTRequest struct {
	db.Service
	NewCategory string
	Categories  []db.Category
}
