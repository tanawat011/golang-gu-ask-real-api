package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BsonD ...
func BsonD(data interface{}, receiver *bson.D) {
	newData := StructToMap(data)
	temp := bson.D{}
	for k, v := range newData {
		temp = append(temp, primitive.E{k, v})
	}
	*receiver = temp
}
