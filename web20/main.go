package main

import (
	"github.com/urfave/negroni"
	"making-web-project-by-go/web20/app"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("web16/public")))
	n.UseHandler(m)

	http.ListenAndServe(":3000", n)
}
