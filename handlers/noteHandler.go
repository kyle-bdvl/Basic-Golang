package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"CrudApp/pkg/inference"
	"CrudApp/tutorial"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

// dependency for the system.
type NoteHandler struct {
	Queries *tutorial.Queries
}

// creating the titleResult struct
type titleResult struct {
	Title string
	Err   error
}

// To create the API endpoint for note creation.
func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {

	input, err := inference.NewNote(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := h.Queries.CreateNote(r.Context(), tutorial.CreateNoteParams{
		Userid:            input.UserID,
		DescriptionOfNote: pgtype.Text{String: input.DescriptionOfNote, Valid: true},
	})

	if err != nil {
		if strings.Contains(err.Error(), "foreign key") {
			http.Error(w, "user id does not exist", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Run AI inference concurrently in the background
	go func(noteID int64, description string) {
		// Note: We use context.Background() here because r.Context() cancels
		// as soon as the HTTP request finishes!
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Best practice note: It's better to load .env in your main.go once on startup,
		// rather than opening the file on every single request.
		if err := godotenv.Load("../../.env"); err != nil {
			log.Println("Error loading .env file")
		}
		apiKey := os.Getenv("auth_secret")

		// Call AI
		title, err := inference.GenerateTitleFromAI(ctx, description, apiKey)
		if err != nil {
			// We must use log here, NOT http.Error(w, ...), because the response
			// has likely already been sent to the user.
			log.Printf("failed to generate title for note %d: %v\n", noteID, err)
			return
		}

		// Update DB with generated title
		err = h.Queries.UpdateTitle(ctx, tutorial.UpdateTitleParams{
			ID: noteID,
			Title: pgtype.Text{
				String: title,
				Valid:  true,
			},
		})

		if err != nil {
			log.Printf("failed to save title for note %d: %v\n", noteID, err)
		}
	}(note.ID, input.DescriptionOfNote)

	// 4. Return the newly created note (without waiting for the AI title)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// getting notes by user
func (h *NoteHandler) ListNotesByUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	notesByUser, err := h.Queries.ListNotesByUser(r.Context(), id)
	if err != nil {
		http.Error(w, "user author not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notesByUser)
}

// Run AI inference concurrently
