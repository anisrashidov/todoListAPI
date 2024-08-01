package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HomeSchema struct {
	Tasks []primitive.M
	Date  string
}
