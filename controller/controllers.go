package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/anisrashidov/todoAPP/model"
	"github.com/anisrashidov/todoAPP/schema"
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

func Index(w http.ResponseWriter, r *http.Request) {
	defer GetAllTasks(w, r)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	params := mux.Vars(r)
	task := getOneTask(task_coll, params["id"])
	json.NewEncoder(w).Encode(task)
	// fmt.Println(task)
	fmt.Println("_______________________________________________________")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	fmt.Println("Get functionality!")
	tasks := getAllTasks(task_coll)
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.ParseFiles(path.Join(cwd, "router", "static", "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	// json.NewEncoder(w).Encode(tasks)
	tmpl.Execute(w, schema.HomeSchema{Tasks: tasks, Date: time.Now().Format("Wed 2017-09-2")})
	// fmt.Println(tasks)
	fmt.Println("_______________________________________________________")
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	fmt.Println("Create functionality!")
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Fatal(err)
		return
	}
	status_code := createTask(task_coll, task)
	w.WriteHeader(status_code)
	fmt.Println("_______________________________________________________")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	fmt.Println("Update functionality!")
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.WriteHeader(updateTask(task_coll, task))
	fmt.Println("_______________________________________________________")
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	var params map[string]interface{}
	fmt.Println("Delete functionality")
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Fatal(err)
	}
	ids_p := params["ids"].(string)
	var ids = strings.Split(ids_p, ",")
	status_code := deleteTasks(task_coll, ids)
	w.WriteHeader(status_code)
	fmt.Println("_______________________________________________________")
}
