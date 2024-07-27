package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anisrashidov/todoAPP/model"
	"github.com/gorilla/mux"
	"github.com/lpernett/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var task_coll *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DB_URL := os.Getenv("DB_URL")
	DB_NAME := os.Getenv("DB_NAME")
	COL_NAME := os.Getenv("COL_NAME")
	clientOption := options.Client().ApplyURI(DB_URL)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("MongoDB connection success")

	task_coll = client.Database(DB_NAME).Collection(COL_NAME)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	task := getOneTask(task_coll, params["id"])
	json.NewEncoder(w).Encode(task)
	fmt.Println(task)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	tasks := getAllTasks(task_coll)
	json.NewEncoder(w).Encode(tasks)
	fmt.Println(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	var task model.Task
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("The following task will be addded: ", task.DueDate)
	status_code := createTask(task_coll, task)
	w.WriteHeader(status_code)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var params map[string]interface{}
	fmt.Println("askdas: ", r.Body)
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.WriteHeader(updateTask(task_coll, params["task_id"].(string)))
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	// var ids_string string
	fmt.Println(r.Body)
	// err := json.NewDecoder(r.Body).Decode(&ids_string)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// var ids = strings.SplitAfter(ids_string, ",")
	// status_code := deleteTasks(task_coll, ids)
	// w.WriteHeader(status_code)
}
