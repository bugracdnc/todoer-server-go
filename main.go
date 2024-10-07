package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
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

func main() {

	print_logo()

	basePath := "/api/v1/todoer"
	port := ":8654"

	id, _ := uuid.NewUUID()
	todoRepository := TodoRepository{todo: []Todo{{Id: id, Todo: "Connect to database", Done: false, CreatedDate: time.Now(), UpdateDate: time.Now()}}}
	todoService := TodoService{todoRepository: todoRepository}
	todoController := TodoController{todoService: todoService}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get(basePath, todoController.getHandler)
	r.Post(basePath, todoController.createHandler)
	r.Get(basePath+"/{id}", todoController.getByIdHandler)
	r.Put(basePath+"/{id}", todoController.updateHandler)
	r.Delete(basePath+"/id", todoController.deleteHandler)

	fmt.Printf("Server listening on http://127.0.0.1%s%s\n", port, basePath)
	http.ListenAndServe(port, r)
}
