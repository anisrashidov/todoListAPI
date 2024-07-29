package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/anisrashidov/todoAPP/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getAllTasks(task_coll *mongo.Collection) []primitive.M {
	var tasks_inq []primitive.M
	cur, err := task_coll.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var task_inq bson.M
		err := cur.Decode(&task_inq)
		if err != nil {
			log.Fatal(err)
		}
		tasks_inq = append(tasks_inq, task_inq)
	}
	return tasks_inq
}

func getOneTask(task_coll *mongo.Collection, task_id string) primitive.M {
	id, err := primitive.ObjectIDFromHex(task_id)
	if err != nil {
		log.Fatal(err)
		return primitive.M{}
	}
	var task_inq bson.M
	task_coll.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task_inq)
	return task_inq
}

func createTask(task_coll *mongo.Collection, task model.Task) int {
	inserted, err := task_coll.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
		return 404
	}
	fmt.Printf("The task was inserted into DB with task ID %d\n", inserted.InsertedID)
	return 200
}

func updateTask(task_coll *mongo.Collection, task_id string) (status_code int) {
	status_code = 200
	id, err := primitive.ObjectIDFromHex(task_id)
	if err != nil {
		log.Fatal(err)
		return 400
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"task_mark": true}}
	res, err := task_coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return 400
	}
	fmt.Println("The updated task looks like: ", res)
	return status_code
}

func deleteTasks(task_coll *mongo.Collection, task_ids []string) int {
	fmt.Println("Here1")
	for ind, task_id := range task_ids {
		fmt.Printf("Here%d\n", ind+2)
		id, err := primitive.ObjectIDFromHex(task_id)
		if err != nil {
			log.Fatal(err)
			return 400
		}
		_, err = task_coll.DeleteOne(context.Background(), bson.M{"_id": id})
		if err != nil {
			log.Fatal(err)
			return 400

		}
		fmt.Printf("The %dth deleted task with task ID %s was deleted!\n", ind, task_id)
	}
	return 200
}
