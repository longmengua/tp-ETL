package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(
	host string,
	port string,
	db string,
	user string,
	password string,
	collection string,
) (
	mCollection *mongo.Collection,
	insert func(data []interface{}),
) {
	var connectionString string
	if port != "" {
		connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
	} else {
		connectionString = fmt.Sprintf("mongodb://%s:%s@%s/%s?retryWrites=true&w=majority", user, password, host, db)
	}
	mcient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	d := mcient.Database(db)
	mCollection = d.Collection(collection)
	insert = func(data []interface{}) {
		result, err := mCollection.InsertMany(context.TODO(), data)
		if err != nil {
			fmt.Printf("Documents insertion error: %v\n", err)
		}
		fmt.Printf("Documents inserted: %v\n", len(result.InsertedIDs))
	}
	return
}
