package db

import (
    "context"
    "fmt"
    "log"
    Bson "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connect() {
   

    Client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://rakesh:10980000gG#@cluster0.sy3jf.mongodb.net/USERS?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer Client.Disconnect(ctx)
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	databases, err := Client.ListDatabaseNames(ctx, Bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	fmt.Println("Connected to MongoDB!")


}


