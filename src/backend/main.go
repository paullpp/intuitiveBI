package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
		table, err := getDBData(body)
		if err != nil {
			fmt.Println("== Error in getDBData:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		tables := table.Results
		connections[numConnections] = Connection{Name: body, Tables: tables}
		numConnections++

		render.Status(r, 201)
	})

	r.Get("/connections/{conn}/tables/{table}/preview", func(w http.ResponseWriter, r *http.Request) {
		idx, err := strconv.Atoi(chi.URLParam(r, "conn"))
		if err != nil {
			fmt.Println("== Error in /connections/:conn/tables str conversion:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		conn := connections[idx]
		schema := strings.Split(conn.Name, "/")
		table := string(schema[len(schema)-1]) + "." + chi.URLParam(r, "table")

		db, err := connectToDb(conn.Name)
		if err != nil {
			fmt.Println("== Unable to establish DB conn:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		query := fmt.Sprintf("select * from %s limit 10;", table)
		fmt.Println("=== Query: ", query)
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println("== Unable to establish DB conn:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		result := []map[string]interface{}{}
		for rows.Next() {
			err := rows.Scan(valuePtrs...)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			rowData := make(map[string]interface{})
			for i, col := range columns {
				val := values[i]
				rowData[col] = fmt.Sprintf("%s", val) // coerce to string
			}

			result = append(result, rowData)
		}
		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Println("== Result: ", result)
		render.JSON(w, r, result)
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
func getDBData(connStr string) (QueryData, error) {
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

func connectToDb(connStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, errors.New("couldn't connect to DB")
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
