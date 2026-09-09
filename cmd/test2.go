package main

import (
	"CrudApp/pkg/inference"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const description = "from pages 16, 18, 19, and 22 of your textbook. Make sure that all questions on these pages are attempted carefully and that your working steps are shown clearly wherever required. Neat presentation, correct calculations, and proper use of mathematical methods will help ensure that your work is complete and easy to review."

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("auth_secret")
	fmt.Printf("API KEY LENGTH: %d\n", len(apiKey))
	title, err := inference.GenerateTitleFromAI(context.Background(), description, apiKey)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title Generated : ", title)
}
