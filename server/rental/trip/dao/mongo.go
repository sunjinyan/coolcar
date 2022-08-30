package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	"coolcar/shared/mgoutil"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tripField = "trip"
	accountIdField = tripField+".accuntId"
)

type Mongo struct {
	col  *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:      db.Collection("trip"),
	}
}

type TripRecord struct {
	mgoutil.IDField
	mgoutil.UpdateAtField
	Trip *rentalpb.Trip `bson:"trip"`
}


func (m *Mongo)CreateTrip(ctx context.Context,trip *rentalpb.Trip) (*TripRecord,error) {

	r := &TripRecord{
		Trip:          trip,
	}
	r.ID = mgoutil.NewObjId()
	r.UpdateAt = mgoutil.UpdateAt()

	_, err := m.col.InsertOne(ctx,r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (m *Mongo)GetTrip(ctx context.Context,tripId  id.TripID,accountId id.AccountID) (*TripRecord,error) {

	objId, err := primitive.ObjectIDFromHex(tripId.String())
	if  err!= nil {
		return nil, err
	}
	res := m.col.FindOne(ctx, bson.M{
		mgoutil.IDFieldName: objId,
		accountIdField:     accountId,
	})
cmp.Diff()
	if err := res.Err(); err != nil {
		return nil, err
	}

	var t  TripRecord //已经分配了空间

	err = res.Decode(&t)
	if err := res.Err(); err != nil {
		return nil, err
	}
	return &t,nil
}