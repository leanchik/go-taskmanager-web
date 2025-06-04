package main

import (
	"net/http"
	"strconv"
	"taskmanagerweb/storage"
	"time"
)

func addTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		title := r.FormValue("title")

		tasks, _ := store.LoadTasks()
		newID := 1
		if len(tasks) > 0 {
			newID = tasks[len(tasks)-1].ID + 1
		}

		newTask := storage.Task{
			ID:        newID,
			Title:     title,
			Done:      false,
			CreatedAt: time.Now(),
		}

		tasks = append(tasks, newTask)
		store.SaveTasks(tasks)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func markTaskDone(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		idStr := r.FormValue("id")
		id, _ := strconv.Atoi(idStr)

		tasks, _ := store.LoadTasks()
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Done = true
				break
			}
		}
		store.SaveTasks(tasks)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		idStr := r.FormValue("id")
		id, _ := strconv.Atoi(idStr)

		tasks, _ := store.LoadTasks()
		var updatedTasks []storage.Task
		for _, task := range tasks {
			if task.ID != id {
				updatedTasks = append(updatedTasks, task)
			}
		}
		store.SaveTasks(updatedTasks)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		idStr := r.FormValue("id")
		newTitle := r.FormValue("title")
		id, _ := strconv.Atoi(idStr)

		tasks, _ := store.LoadTasks()
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Title = newTitle
				break
			}
		}
		store.SaveTasks(tasks)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
