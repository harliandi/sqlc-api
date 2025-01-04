package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"harliandi.dev/sqlc-api/author"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	log.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	q := author.New(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/author", func(w http.ResponseWriter, r *http.Request) {
		authors, err := q.ListAuthors(context.Background())
		if err != nil {
			render.JSON(w, r, err)
			return
		}

		render.JSON(w, r, authors)
	})

	r.Get("/author/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			render.JSON(w, r, err)
			return
		}
		authors, err := q.GetAuthor(context.Background(), int64(id))
		if err != nil {
			render.JSON(w, r, err)
			return
		}

		render.JSON(w, r, authors)
	})

	r.Post("/author", func(w http.ResponseWriter, r *http.Request) {
		type createAuthorPayload struct {
			Name string `json:"name"`
			Bio  string `json:"bio"`
		}

		var payloadAuthor createAuthorPayload
		if err := json.NewDecoder(r.Body).Decode(&payloadAuthor); err != nil {
			render.JSON(w, r, err)
			return
		}
		defer r.Body.Close()

		authors, err := q.CreateAuthor(context.Background(), author.CreateAuthorParams{Name: payloadAuthor.Name, Bio: sql.NullString{String: payloadAuthor.Bio, Valid: true}})
		if err != nil {
			render.JSON(w, r, err)
			return
		}
		data, err := json.Marshal(authors)
		if err != nil {
			render.JSON(w, r, err)
			return
		}

		render.JSON(w, r, data)
	})

	r.Delete("/author/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			render.JSON(w, r, err)
			return
		}

		err = q.DeleteAuthor(context.Background(), int64(id))
		if err != nil {
			render.JSON(w, r, err)
			return
		}

		render.JSON(w, r, fmt.Sprint("success delete author with id: ", id))
	})

	serverPort := os.Getenv("SERVER_PORT")
	log.Println("server running on port:", serverPort)
	http.ListenAndServe(":"+serverPort, r)
}
