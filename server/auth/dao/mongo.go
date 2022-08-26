package dao

import (
	"context"
	mgo "coolcar/shared"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	col *mongo.Collection
}

const openIDField = "open_id"

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

func (m *Mongo)ResolveAccountID(ctx context.Context,openID string) (ObjectID string,err error) {

	var Upsert bool = true
	var ReturnDocument options.ReturnDocument = options.After
	res := m.col.FindOneAndUpdate(ctx, bson.M{
		openIDField: openID,
	},
	mgo.Set(bson.M{
			openIDField: openID,
	}),
	&options.FindOneAndUpdateOptions{
		ReturnDocument: &ReturnDocument,
		Upsert:         &Upsert,
	})
	if err := res.Err(); err != nil {
		return "",fmt.Errorf("cannot find one and update: %v",err)
	}
	var row mgo.ObjID
	err = res.Decode(&row)
	if err != nil {
		return "",fmt.Errorf("cannot find one and update: %v",err)
	}

	return row.ID.Hex(),nil
}