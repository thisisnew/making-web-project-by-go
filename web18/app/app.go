package app

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"strconv"
	"time"
)

var rd *render.Render

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	id := len(todoMap) + 1
	todo := &Todo{
		ID:        id,
		Name:      name,
		Completed: false,
		CreatedAt: time.Time{},
	}
	todoMap[id] = todo
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}

}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"

	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}

}

func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/todos", getTodoListHandler).Methods(http.MethodGet)
	r.HandleFunc("/todos", addTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods(http.MethodDelete)
	r.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/", indexHandler)

	return r
}
