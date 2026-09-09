package inference

import (
	"encoding/json"
	"errors"
	"net/http"
)

type NoteInput struct {
	// Fields MUST start with a capital letter for JSON decoding to work
	DescriptionOfNote string `json:"description"`
	UserID            int64  `json:"userId"`
}

// NewNoteInput is a package-level function (not a method)
func NewNote(r *http.Request) (*NoteInput, error) {
	input := NoteInput{}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, errors.New("invalid request body")
	}

	if input.DescriptionOfNote == "" {
		return nil, errors.New("description is required")
	}
	if input.UserID == 0 {
		return nil, errors.New("user id is required")
	}

	return &input, nil
}
