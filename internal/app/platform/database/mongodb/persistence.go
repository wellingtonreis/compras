package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var schema = GetMongoSchema()

func (db *MongoDB) InsertData(collectionName string, data []interface{}) error {
	collection := db.Client.Database(schema).Collection(collectionName)
	_, err := collection.InsertMany(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) GetData(collectionName string, ctx context.Context, filter interface{}, reference interface{}) error {
	collection := db.Client.Database(schema).Collection(collectionName)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, reference); err != nil {
		return err
	}

	return nil
}

func (db *MongoDB) UpdateData(collectionName string, filter []byte, update []byte) error {
	collection := db.Client.Database(schema).Collection(collectionName)
	_, err := collection.UpdateMany(context.Background(), filter, update)
	return err
}

func (db *MongoDB) DeleteData(collectionName string, filter []byte) error {
	collection := db.Client.Database(schema).Collection(collectionName)
	_, err := collection.DeleteMany(context.Background(), filter)
	return err
}

type Counter struct {
	ID    string `bson:"_id"`
	Value int64  `bson:"value"`
}

func (db *MongoDB) GetNextSequenceValue(sequenceName string) (int64, error) {
	countersCollection := db.Client.Database(schema).Collection("sequences")

	filter := bson.M{"_id": sequenceName}
	update := bson.M{"$inc": bson.M{"value": 1}}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)

	var counter Counter
	err := countersCollection.FindOneAndUpdate(context.TODO(), filter, update, options).Decode(&counter)
	if err != nil {
		return 0, err
	}

	return counter.Value, nil
}
