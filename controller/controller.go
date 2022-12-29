package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://shailendra:<password>@cluster0.gud6nr5.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongodb
func init() {

	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Connection failed")
	}
	fmt.Println("Connection succeeded")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")

}
