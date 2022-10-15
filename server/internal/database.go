package internal

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client     *mongo.Client
	Disconnect chan bool
}

func NewDatabase() *Database {
	return &Database{
		Disconnect: make(chan bool),
	}
}

func InitDatabase(db *Database) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' env variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	db.Client = client

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	log.Println("Connected to the database")

	select {
	case <-db.Disconnect:
		log.Println("Closing database")
		return
	}
}
