package main

import (
	"html/template"
	"net/http"
	"task-3/configuration"

	"github.com/sirupsen/logrus"
)

func showName(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		logrus.Fatalln(err)
	}

	data := struct {
		Name string
	}{
		Name: name,
	}

	if err := tmpl.ExecuteTemplate(w, "Task-3", data); err != nil {
		logrus.Fatalln(err)
	}
}

func main() {
	conf := configuration.Config()

	route := http.NewServeMux()
	route.HandleFunc("/", showName)

	logrus.Fatalln(http.ListenAndServe(conf.ADDRESS, route))
}
