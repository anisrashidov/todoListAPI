package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type UnixTime time.Time

// func (t *UnixTime) UnmarshalJSON(b []byte) error {
// 	var timestamp int64
// 	if err := json.Unmarshal(b, &timestamp); err != nil {
// 		return err
// 	}
// 	*t = UnixTime(time.Unix(0, timestamp*int64(time.Millisecond)))
// 	return nil
// }

type Task struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"task_name,omitempty"`
	Description string             `json:"task_desc,omitempty"`
	DueDate     string             `json:"task_due,omitempty"`
	MarkedDone  bool               `json:"task_mark,omitempty"`
}
