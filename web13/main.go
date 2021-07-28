package main

import (
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	log.Print("postMessageHandler", msg, name)
}

func main() {
	mux := pat.New()
	mux.Post("/messages", postMessageHandler)

	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("web13/public")))
	n.UseHandler(mux)
	http.ListenAndServe(":3000", n)
}
