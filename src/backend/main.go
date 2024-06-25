package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type QueryData = struct {
	Results []string `json:"results"`
}

func main() {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	r.Use(cors.Handler)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("== Root")
		render.Status(r, 200)
	})

	r.Post("/connect", func(w http.ResponseWriter, r *http.Request) {
		var body string

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		connectToDb(body)
		render.Status(r, 201)
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func connectToDb(connStr string) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println("Error connecting to DB:", err)
		return
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		fmt.Println("Error fetching tables:", err)
		return
	}
	defer rows.Close()

	var data []string
	for rows.Next() {
		var row string
		err := rows.Scan(&row)
		if err != nil {
			return
		}
		data = append(data, row)
	}
	res := QueryData{Results: data}
	fmt.Println(res)
}
