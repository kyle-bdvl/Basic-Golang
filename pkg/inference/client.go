package inference

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const inferenceURL = "http://192.168.2.23:8080/v1/chat/completions"

type TitleResponse struct {
	Title string `json:"Title"`
}

func GenerateTitleFromAI(ctx context.Context, description string, apiKey string) (string, error) {

	reqBody := ChatCompletionsRequest{
		Model: "GrafanaTest/gpt-oss-20b",

		Messages: []Message{
			{
				Role:    "system",
				Content: "Generate a short note title ONLY MAXIMUM OF THREE WORDS. Respond ONLY using the provided JSON schema.",
			},
			{
				Role:    "user",
				Content: description,
			},
		},

		ResponseFormat: ResponseFormat{
			Type: "json_schema",

			JSONSchema: JSONSchema{
				Name:   "note_title",
				Strict: true,

				Schema: JSONSchemaBody{
					Type: "object",

					Properties: map[string]SchemaProperty{
						"Title": {
							Type:        "string",
							Description: "Generated title",
						},
					},

					Required: []string{
						"Title",
					},

					AdditionalProperties: false,
				},
			},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		inferenceURL,
		bytes.NewReader(body),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("AI server returned %d : %s", resp.StatusCode, string(b))
	}

	var aiResp ChatCompletionsResponse

	if err := json.NewDecoder(resp.Body).Decode(&aiResp); err != nil {
		return "", err
	}

	if len(aiResp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned")
	}

	// Content itself is JSON text
	var title TitleResponse

	if err := json.Unmarshal(
		[]byte(aiResp.Choices[0].Message.Content),
		&title,
	); err != nil {
		return "", err
	}

	return title.Title, nil
}
