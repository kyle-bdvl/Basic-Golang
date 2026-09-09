package handlers

import (
	"CrudApp/pkg/inference"
	"CrudApp/tutorial"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserHandler struct {
	Queries *tutorial.Queries
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	input, err := inference.NewUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Execute the query using pgtype.Text
	user, err := h.Queries.CreateUser(r.Context(), pgtype.Text{
		String: input.Username,
		Valid:  true, // This tells Postgres the value is NOT NULL
	})
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// getting all users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUser, err := h.Queries.GetAllUsers(r.Context())

	if err != nil {
		http.Error(w, "Server is down ?", http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUser)
}

// creating the get requests
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	// ** second use case **
	// idParam := r.URL.Query().Get("id")

	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	//  r.Context() comes from HTTP Requests and uses it to exe. SQL Query
	user, err := h.Queries.GetUser(r.Context(), id)

	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

// another get request
func (h *UserHandler) GetUsersByQueryParam(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(idParam, 10, 32)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.Queries.GetUser(r.Context(), id)

	if err != nil {
		http.Error(w, "user doesn't exist", http.StatusNotFound)
		return
	}
	// handling the id and the return json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
