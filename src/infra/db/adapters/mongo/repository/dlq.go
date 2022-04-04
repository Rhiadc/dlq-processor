package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DLQRecord struct {
	Date      time.Time `bson:"date"`
	Msg       string    `bson:"message"`
	Processed bool      `bson:"processed"`
}

type DLQRepository struct {
	database *mongo.Database
}
type DLQRepo interface {
	//@TODO
	//delete DLQRecord
	//update DLQRecord
}

var (
	DLQCollection = "dlq"
)

func NewDLQRepository(database *mongo.Database) *DLQRepository {
	return &DLQRepository{database: database}
}

func (ref DLQRepository) InsertMessage(m DLQRecord) error {
	coll := ref.database.Collection(DLQCollection)
	_, err := coll.InsertOne(context.TODO(), m)
	if err != nil {
		return err
	}
	return nil
}

func (ref DLQRepository) GetAllMessages() ([]DLQRecord, error) {
	var msg DLQRecord
	var DLQRecords []DLQRecord
	coll := ref.database.Collection(DLQCollection)
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		defer cursor.Close(context.TODO())
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&msg)
		if err != nil {
			return DLQRecords, err
		}
		DLQRecords = append(DLQRecords, msg)
	}
	return DLQRecords, nil
}

func (ref DLQRepository) GetMessagesByDateRange(days int) ([]DLQRecord, error) {
	var msg DLQRecord
	var DLQRecords []DLQRecord
	coll := ref.database.Collection(DLQCollection)
	cursor, err := coll.Find(context.TODO(), bson.M{"date": bson.M{
		"$gte": primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, days))},
		"processed": bson.M{"$eq": false},
	})
	if err != nil {
		fmt.Println(err)
		defer cursor.Close(context.TODO())
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&msg)
		if err != nil {
			return DLQRecords, err
		}
		DLQRecords = append(DLQRecords, msg)
	}
	return DLQRecords, nil
}

func (ref DLQRepository) GetProcessedMessages() ([]DLQRecord, error) {
	var msg DLQRecord
	var DLQRecords []DLQRecord
	coll := ref.database.Collection(DLQCollection)
	cursor, err := coll.Find(context.TODO(), bson.M{
		"processed": bson.M{"$eq": true},
	})
	if err != nil {
		defer cursor.Close(context.TODO())
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&msg)
		if err != nil {
			return DLQRecords, err
		}
		DLQRecords = append(DLQRecords, msg)

	}
	return DLQRecords, nil
}

func (ref DLQRepository) SetProcessedMessage(date time.Time) error {
	coll := ref.database.Collection(DLQCollection)
	filter := bson.D{{"date", date}}
	update := bson.D{{"$set", bson.D{{"processed", true}}}}
	_, err := coll.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
	return err
}

func (ref DLQRepository) GetMessageByDate(date time.Time) (DLQRecord, error) {
	var msg DLQRecord
	coll := ref.database.Collection(DLQCollection)
	err := coll.
		FindOne(context.TODO(), bson.D{{"date", date}}).
		Decode(&msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (ref DLQRepository) DeleteMessageByDate(date time.Time) error {
	coll := ref.database.Collection(DLQCollection)
	_, err := coll.DeleteOne(context.TODO(), bson.D{{"date", primitive.NewDateTimeFromTime(date)}})
	if err != nil {
		return err
	}
	return nil
}
