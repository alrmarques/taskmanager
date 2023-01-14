package routes

import (
	"config"
	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"models"
	"net/http"
)

func Init() {
	router := mux.NewRouter()

	// Routers for task management
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	router.HandleFunc("/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	router
	HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	http.Handle("/", router)

}

//GetAllTasks retrieves all tasks from the database

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	// Code to retrieve all tasks from the database
}

// GetTasks retrieves a specific task from the database

func GetTasks(w http.ResponseWriter, r *http.Request) {
	// Code to retrieve a specific task from the database
}

// CreateTask creates a new task in the database
func CreateTask(w http.ResponseWriter, r *http.Request) {
	// Code to create a new task in the database
}

// UpdateTask updates a specific task in the database
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Code to update a specific task in the database
}

// DeleteTask deletes a specific task in the database
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Code to delete a specific task in the database
}
