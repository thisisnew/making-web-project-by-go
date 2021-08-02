package main

import (
	"making-web-project-by-go/web22/app"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()

	http.ListenAndServe(":3000", m)
}
