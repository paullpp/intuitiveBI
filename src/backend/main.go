package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type QueryData = struct {
	Results []string
}

type Connection = struct {
	Name   string
	Tables []string
}

var connections = make(map[int]Connection)
var numConnections = 0

func main() {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	r.Use(cors.Handler)
	r.Use(middleware.Logger)

	/*
	 * Root Route, currently useless
	 */
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		render.Status(r, 200)
	})

	/*
	 * Returns all connections currently stored in the connections map
	 */
	r.Get("/connections", func(w http.ResponseWriter, r *http.Request) {

		render.JSON(w, r, connections)
	})

	/*
	 * Returns all tables for a given connection
	 */
	r.Get("/connections/{conn}/tables", func(w http.ResponseWriter, r *http.Request) {
		conn, err := strconv.Atoi(chi.URLParam(r, "conn"))
		if err != nil {
			fmt.Println("== Error in /connections/:conn/tables str conversion:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		render.JSON(w, r, connections[conn].Tables)
	})

	/*
	 * Adds a new connection. Stores it in the connections map
	 */
	r.Post("/connections", func(w http.ResponseWriter, r *http.Request) {
		var body string

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Add connection with name and empty arr for tables
		table, err := connectToDb(body)
		if err != nil {
			fmt.Println("== Error in connectToDb:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		tables := table.Results
		connections[numConnections] = Connection{Name: body, Tables: tables}
		numConnections++

		render.Status(r, 201)
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

/*
 * Connects to the DB using the provided connection string
 * Returns an array of all tables in the DB
 */
func connectToDb(connStr string) (QueryData, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return QueryData{}, errors.New("couldn't connect to DB")
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("SHOW TABLES")
	if err != nil {

		return QueryData{}, errors.New("couldn't fetch tables")
	}
	defer rows.Close()

	var data []string
	for rows.Next() {
		var row string
		err := rows.Scan(&row)
		if err != nil {
			return QueryData{}, errors.New("error scanning rows")
		}
		data = append(data, row)
	}
	res := QueryData{Results: data}

	return res, nil
}
