package model

import (
	"context"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connections = make(map[string]*mongo.Client)

// MongoCollection struct
type MongoCollection struct {
}

// NewMongoCollection :: used to init struct
func NewMongoCollection(db string, collection string) *mongo.Collection {
	db = strings.ToUpper(db)
	if _, exist := connections[db]; exist == false {
		log.Printf(`Connection to '%s' doesn't exist, trying to make new connection`,
			db)
		connectionString := os.Getenv(db + "_CONNECTION_STRING")
		if connectionString == "" {
			connectionString = "mongodb://localhost:27017/"
		}
		log.Printf(`%s_CONNECTION_STRING=%s`, db, connectionString)
		clientOptions := options.Client().ApplyURI(connectionString)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		log.Printf(`Connected to '%s'`, db)
		connections[db] = client
	}
	return connections[db].Database(db).Collection(collection)
}
