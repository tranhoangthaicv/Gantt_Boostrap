package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title     string
	WorkShip  string
	Durations int
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("index.html"))
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "1", WorkShip: "User 1", Durations: 1},
				{Title: "2", WorkShip: "User 2", Durations: 2},
				{Title: "3", WorkShip: "User 3", Durations: 3},
				{Title: "4", WorkShip: "User 4", Durations: 4},
				{Title: "5", WorkShip: "User 5", Durations: 5},
				{Title: "6", WorkShip: "User 6", Durations: 6},
				{Title: "7", WorkShip: "User 7", Durations: 8},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
