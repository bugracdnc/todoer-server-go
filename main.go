package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Import the file source driver
	_ "github.com/lib/pq"                                // PostgreSQL driver
)

func print_logo() {
	var logo []string = make([]string, 7)

	logo[0] = "░▒▓████████▓▒░▒▓██████▓▒░░▒▓███████▓▒░ ░▒▓██████▓▒░░▒▓████████▓▒░▒▓███████▓▒░  "
	logo[1] = "   ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ "
	logo[2] = "   ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ "
	logo[3] = "   ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓██████▓▒░ ░▒▓███████▓▒░  "
	logo[4] = "   ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ "
	logo[5] = "   ░▒▓█▓▒░  ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ "
	logo[6] = "   ░▒▓█▓▒░   ░▒▓██████▓▒░░▒▓███████▓▒░ ░▒▓██████▓▒░░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░ "

	fmt.Println()
	for _, line := range logo {
		fmt.Println(line)
	}
	fmt.Println()
	fmt.Println()
}

var dbOp DBOperations

var tokenLength = 4
var bearerLength = 6

func generateTokenAndSaveUser(w http.ResponseWriter, r *http.Request) {
	token, _ := randomHex(tokenLength)
	user := User{Token: token, Name: ""}
	dbOp.saveUserToDb(user)
	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(token)
	if encodeErr != nil {
		fmt.Printf("error (generateToken): %s", encodeErr.Error())
		http.Error(w, "Error!", http.StatusSeeOther)
	}
}

func validateTokenFromUser(r *http.Request) bool {
	users, err := dbOp.getUsersFromDb()
	if err != nil {
		log.Fatalf("Failed to get users from db: %v", err)
		return false
	}
	bearerToken := r.Header.Get("Authorization")
	if len(bearerToken) < ((tokenLength * 2) + bearerLength + 1) { //+1 is for the space
		fmt.Println("Bad Bearer Token")
		return false
	}
	reqToken := strings.Split(bearerToken, " ")[1]
	for _, user := range users {
		if user.Token == reqToken {
			return true
		}
	}
	fmt.Println("Failed to authenticate with all tokens in db")
	return false
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func main() {

	print_logo()

	databaseSource := "postgres://bugra:123789bugra@localhost:5432/todoerdb?sslmode=disable"

	db, err := sql.Open("postgres", databaseSource)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	dbOp = DBOperations{DB: db}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not start SQL driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("Could not start migration: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	basePath := "/api/v1/todoer"
	port := ":8654"

	todoRepository := TodoRepository{todos: []Todo{}, DB: dbOp}
	todoService := TodoService{todoRepository: todoRepository}
	todoController := TodoController{todoService: todoService}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get(basePath+"/login", generateTokenAndSaveUser)

	r.Get(basePath, todoController.getHandler)
	r.Post(basePath, todoController.createHandler)
	r.Get(basePath+"/{id}", todoController.getByIdHandler)
	r.Put(basePath+"/{id}", todoController.updateHandler)
	r.Delete(basePath+"/id", todoController.deleteHandler)

	fmt.Printf("Server listening on http://127.0.0.1%s%s\n", port, basePath)
	http.ListenAndServe(port, r)
}
