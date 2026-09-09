package main

import (
	"CrudApp/handlers"
	"CrudApp/tutorial"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// exporting dotenv
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbEnv := os.Getenv("GOOSE_DBSTRING")
	// its like a phone number to check for the pool to check the connection
	ctx := context.Background()
	// generating the pool connection (just like in node )
	pool, err := pgxpool.New(ctx, dbEnv)
	if err != nil {
		log.Fatal(err)
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("database ping failed:", err)
	} else {
		fmt.Println("There is a ping")
	}

	// Happens once main function finishes. Good for clean up
	defer pool.Close()
	queries := tutorial.New(pool)
	userHandler := &handlers.UserHandler{
		Queries: queries,
	}
	noteHandler := &handlers.NoteHandler{
		Queries: queries,
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("welcome to my go server"))
		fmt.Fprintln(w, "Sends text back, this server works !")
	})
	r.Post("/createUser", userHandler.CreateUser)
	r.Get("/getUser/{id}", userHandler.GetUsers)
	r.Get("/getAllUsers", userHandler.GetAllUsers)
	r.Get("/getUsersByQueryParam", userHandler.GetUsersByQueryParam)
	// Requests for Notes
	r.Post("/createNote", noteHandler.CreateNote)
	r.Get("/getNoteByUser/{id}", noteHandler.ListNotesByUser)
	http.ListenAndServe(":8181", r)
}
