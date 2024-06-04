package main

import (
	"log"
	"net/http"
	"text/template"
)

type Todo struct {
	Id int
	Message string
}

func main() {

	data := map[string][]Todo{
		"ToDos": {
			Todo{Id: 1, Message: "Buy Milk"},
		},
	}

	todoHandler := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("todo.html"));

		template.Execute(w, data);
	}

	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		template := template.Must(template.ParseFiles("todo.html"));
		todo := Todo{ Id: len(data["ToDos"]) + 1, Message: message }

		template.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", todoHandler);

	http.HandleFunc("/add-todo", addTodoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil));
}