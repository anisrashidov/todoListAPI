package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/anisrashidov/todoAPP/model"
	"github.com/gorilla/mux"
	"github.com/lpernett/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fmt.Println(task)
	fmt.Println("_______________________________________________________")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	tasks := getAllTasks(task_coll)
	tmpl, err := template.ParseFiles("../todoListAPI/router/static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	// json.NewEncoder(w).Encode(tasks)
	tmpl.Execute(w, map[string]([]primitive.M){"Tasks": tasks})
	fmt.Println(tasks)
	fmt.Println("_______________________________________________________")
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("The following task will be addded: ", task)
	status_code := createTask(task_coll, task)
	w.WriteHeader(status_code)
	fmt.Println("_______________________________________________________")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var params map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(params["task_id"].(string))
	w.WriteHeader(updateTask(task_coll, params["task_id"].(string)))
	fmt.Println("_______________________________________________________")
}

func DeleteTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	var params map[string]interface{}
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Fatal(err)
	}
	ids_p := params["ids"].(string)
	fmt.Println(ids_p)
	var ids = strings.Split(ids_p, ",")
	fmt.Println(ids)
	status_code := deleteTasks(task_coll, ids)
	w.WriteHeader(status_code)
	fmt.Println("_______________________________________________________")
}
