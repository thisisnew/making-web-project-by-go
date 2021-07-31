package main

import (
	"github.com/urfave/negroni"
	"making-web-project-by-go/web16/app"
	"net/http"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":3000", n)
}
