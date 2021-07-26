package main

import (
	"making-web-project-by-go/web7/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
