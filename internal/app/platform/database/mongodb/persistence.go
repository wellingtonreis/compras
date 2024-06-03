package mongodb

import (
	"context"
)

func (db *MongoDB) InsertData(schema string, collectionName string, data []interface{}) error {
	collection := db.Client.Database(schema).Collection(collectionName)
	_, err := collection.InsertMany(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) GetData(schema string, collectionName string) ([]interface{}, error) {
	var items []interface{}

	collection := db.Client.Database(schema).Collection(collectionName)
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var item interface{}
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (db *MongoDB) UpdateData(schema string, collectionName string, filter []byte, update []byte) error {
	collection := db.Client.Database(schema).Collection(collectionName)
	_, err := collection.UpdateMany(context.Background(), filter, update)
	return err
}

func (db *MongoDB) DeleteData(schema string, collectionName string, filter []byte) error {
	collection := db.Client.Database(schema).Collection(collectionName)
	_, err := collection.DeleteMany(context.Background(), filter)
	return err
}
