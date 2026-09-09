package inference

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type UserInput struct {
	Username string `json:"username"`
}

func NewUser(r *http.Request) (*UserInput, error) {
	var input UserInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, errors.New("invalid request body")
	}
	if strings.TrimSpace(input.Username) == "" {
		return nil, errors.New("Username is required")
	}

	return &input, nil

}
