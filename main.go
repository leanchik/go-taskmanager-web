package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"taskmanagerweb/storage"
)

var store storage.Storage

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tasks, _ := store.LoadTasks()

	var doneTasks, noteDoneTasks []storage.Task
	for _, task := range tasks {
		if task.Done {
			doneTasks = append(doneTasks, task)
		} else {
			noteDoneTasks = append(noteDoneTasks, task)
		}
	}
	sortedTasks := append(doneTasks, noteDoneTasks...)

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, sortedTasks)
}

func main() {
	storageType := os.Getenv("STORAGE")

	if storageType == "postgres" {
		connStr := "host=localhost port=5432 user=postgres password=lean4real dbname=postgres sslmode=disable"
		var err error
		store, err = storage.NewPostgresStorage(connStr)
		if err != nil {
			log.Fatal("Ошибка подключения к базе:", err)
		}
		log.Println("Хранилище: PostgreSQL")
	} else {
		store = &storage.JSONStorage{FileName: "handlers.json"}
		log.Println("Хранилище: JSON")
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addTask)
	http.HandleFunc("/done", markTaskDone)
	http.HandleFunc("/delete", deleteTask)
	http.HandleFunc("/edit", editTask)
	log.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
