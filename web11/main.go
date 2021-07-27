package main

import (
	"os"
	"text/template"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{
		Name:  "thisisnew",
		Email: "thisisnew@naver.com",
		Age:   34,
	}

	user2 := User{
		Name:  "thisisnew2",
		Email: "thisisnew2@naver.com",
		Age:   14,
	}

	users := []User{user, user2}

	tmpl, err := template.New("Tmpl1").ParseFiles("web11/templates/tmpl1.tmpl", "web11/templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
