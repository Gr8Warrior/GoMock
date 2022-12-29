package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gr8warrior/mongomock/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//MONGODB helpers - file

//insert 1 record

func insertOneMovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal("Insert to mongo failed")
	}

	fmt.Println("Inserted one movie with id ", inserted.InsertedID)

}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal("Update failed")
	}

	fmt.Println("Modified count ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal("Delete failed")
	}

	fmt.Println("Delete count ", result.DeletedCount)
}

func deleteAllMovie() int {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal("Delete failed")
	}

	fmt.Println("Delete count ", deleteResult.DeletedCount)

	return int(deleteResult.DeletedCount)

}

func getAllMovies() []primitive.M {

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal("Unable to find movies")
	}

	defer cur.Close(context.Background())

	var movies []primitive.M
	for cur.Next(context.Background()) {

		var movie bson.M

		err := cur.Decode(&movie)

		if err != nil {
			log.Fatal("Decode Failed")
		}
		movies = append(movies, movie)

	}

	return movies
}

//Actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	err := json.NewEncoder(w).Encode(allMovies)
	if err != nil {
		log.Fatal("Unable to fetch movies")
	} else {
		fmt.Println("Response sent")
	}

}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add movie")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println("Add movie")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()

	json.NewEncoder(w).Encode(count)
}

func ServerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Shailu")
	w.Write([]byte("<h1> Welcome to Mock API Server1 </h1>"))
}
