package main

import (
	"database/sql"
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
		// enter db creds here: user:password@tcp(host:port)/db_name
		db, err := sql.Open("mysql", "user:password@/db")
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

		rows, err := db.Query("SHOW TABLES")
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer rows.Close()

		var data []string
		for rows.Next() {
			var row string
			err := rows.Scan(&row)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data = append(data, row)
		}
		res := QueryData{Results: data}
		fmt.Println(res)
		render.JSON(w, r, res)
	})

	http.ListenAndServe(":3000", r)
}
