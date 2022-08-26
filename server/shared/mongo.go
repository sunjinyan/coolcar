package mgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


//都可能是公用的
const IDField = "_id"

type ObjID struct {
	ID primitive.ObjectID `bson:"_id"`
}

//Set return a $set update document
func Set(v interface{}) bson.M {
	//primitive.ObjectIDFromHex()
	return bson.M{
		"$set": v,
	}
}
