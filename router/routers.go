package router

import (
	"github.com/anisrashidov/todoAPP/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", controller.GetAllTasks).Methods("GET")
	router.HandleFunc("/task/{id}", controller.GetTask).Methods("GET")
	router.HandleFunc("/delete_tasks", controller.DeleteTasks).Methods("DELETE")
	router.HandleFunc("/create", controller.CreateTask).Methods("PUT")
	router.HandleFunc("/edit", controller.UpdateTask).Methods("POST")
	return router
}
