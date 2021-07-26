package myapp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

var usersMap map[int]*User
var lastID int

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUser struct {
	ID               int       `json:"id"`
	UpdatedFirstName bool      `json:"updated_first_name"`
	FirstName        string    `json:"first_name"`
	UpdatedLastName  bool      `json:"updated_last_name"`
	LastName         string    `json:"last_name"`
	UpdatedEmail     bool      `json:"updated_email"`
	Email            string    `json:"email"`
	CreatedAt        time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if len(usersMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}

	users := []*User{}
	for _, u := range usersMap {
		users = append(users, u)
	}
	data, _ := json.Marshal(users)
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}
func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user, ok := usersMap[id]

	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	usersMap[user.ID] = user

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	updateUser := &User{}
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user, ok := usersMap[updateUser.ID]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:", updateUser.ID)
		return
	}

	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}
	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	_, ok := usersMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:", id)
		return
	}

	delete(usersMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User ID:", id)
}

func NewHttpHandler() http.Handler {
	usersMap = make(map[int]*User)
	lastID = 0

	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	mux.HandleFunc("/users", usersHandler).Methods(http.MethodGet)
	mux.HandleFunc("/users", createUserHandler).Methods(http.MethodPost)
	mux.HandleFunc("/users", updateUserHandler).Methods(http.MethodPut)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler).Methods(http.MethodGet)
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods(http.MethodDelete)

	return mux
}
