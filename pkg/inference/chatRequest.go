package inference

type ChatCompletionsRequest struct {
	Model          string         `json:"model"`
	Messages       []Message      `json:"messages"`
	ResponseFormat ResponseFormat `json:"response_format"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseFormat struct {
	Type       string     `json:"type"`
	JSONSchema JSONSchema `json:"json_schema"`
}

type JSONSchema struct {
	Name   string         `json:"name"`
	Strict bool           `json:"strict"`
	Schema JSONSchemaBody `json:"schema"`
}

type JSONSchemaBody struct {
	Type                 string                    `json:"type"`
	Properties           map[string]SchemaProperty `json:"properties"`
	Required             []string                  `json:"required"`
	AdditionalProperties bool                      `json:"additionalProperties"`
}

type SchemaProperty struct {
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

// Example usage on how to use this inference
/*
req := ChatCompletionsRequest{
	Model: "Test/Qwen3.6-27B",
	Messages: []Message{
		{
			Role:    "system",
			Content: "You are a helpful data extraction assistant.",
		},
		{
			Role:    "user",
			Content: "1.eggs, 2. Milk, 3.Bananas, 4.watermelon 5. Haribbo GummyBear",
		},
	},
	ResponseFormat: ResponseFormat{
		Type: "json_schema",
		JSONSchema: JSONSchema{
			Name:   "Title of Notes",
			Strict: true,
			Schema: JSONSchemaBody{
				Type: "object",
				Properties: map[string]SchemaProperty{
					"Title": {
						Type:        "string",
						Description: "The name of the Title based on the description of the note",
					},
				},
				Required:             []string{"Title"},
				AdditionalProperties: false,
			},
		},
	},
}
*/
