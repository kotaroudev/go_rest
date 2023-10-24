package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kotaroudev/go_rest/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

// Las anotaciones tipo `json:"key"`
// son como se va tratar las propiedades del struct
// una vez serializadas en el json.
// Message es en go; "message" es en el json

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to Platzi Go",
			Status:  true,
		})
	}
}
