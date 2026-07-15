package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(
		"SELECT id, name, email FROM users",
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	var users []User

	for rows.Next() {

		var user User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}
func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.QueryRow(
		"INSERT INTO users(name, email) VALUES($1, $2) RETURNING id",
		user.Name,
		user.Email,
	).Scan(&user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {

	idString := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var user User

	err = db.QueryRow(
		"SELECT id, name, email FROM users WHERE id=$1",
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		),

	)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL")

	http.HandleFunc("/users/", getUserByID)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		getUsers(w, r)

	case http.MethodPost:
		createUser(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
})

	fmt.Println("Server running on port 8991")

	log.Fatal(
		http.ListenAndServe(":8991", nil),
	)
}
