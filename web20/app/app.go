package app

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"making-web-project-by-go/web20/model"
	"net/http"
	"strconv"
)

var rd *render.Render = render.New()

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := a.db.AddTodo(name)

	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)

	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}
}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := a.db.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}
}

func (a *AppHandler) Close() {
	a.db.Close()
}

func MakeHandler() *AppHandler {
	r := mux.NewRouter()

	a := &AppHandler{
		Handler: r,
		db:      model.NewDbHandler(),
	}

	r.HandleFunc("/todos", a.getTodoListHandler).Methods(http.MethodGet)
	r.HandleFunc("/todos", a.addTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods(http.MethodDelete)
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods(http.MethodGet)
	r.HandleFunc("/", a.indexHandler)

	return a
}
