package controllers

import (
	"context"
	"corses/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "moviesVerse"
const collectionList = "movieList"

var collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dt := os.Getenv("MONGODB_URI")
	if dt != "" {
		connectionString := dt
		clientOption := options.Client().ApplyURI(connectionString)
		client, err := mongo.Connect(context.TODO(), clientOption)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("MongoDB connection success")
		collection = client.Database(dbName).Collection(collectionList)
		fmt.Println("collection instance is ready")
	} else {
		log.Fatal("error with env")
	}
}

func getMovieById(movieId string) (entity.MovieEntity, error) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return entity.MovieEntity{}, fmt.Errorf("invalid movie ID: %v", err)
	}

	filter := bson.M{"_id": id}
	var movie entity.MovieEntity

	err = collection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		return entity.MovieEntity{}, fmt.Errorf("error finding document: %v", err)
	}

	return movie, nil
}
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

func deleteMovieById(movieId string) int64 {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteCount.DeletedCount
}

func deleteAllMovies() int64 {
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NUmber of movies delete: ", deleted.DeletedCount)
	return deleted.DeletedCount
}

func addMovies(movies entity.MovieEntity) {
	inserted, err := collection.InsertOne(context.Background(), movies)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted movie", inserted.InsertedID)

}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie entity.MovieEntity
	_ = json.NewDecoder(r.Body).Decode(&movie)
	fmt.Print(movie)
	addMovies(movie)
	json.NewEncoder(w).Encode(movie)
}

func GetMoviesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	movie, err := getMovieById(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	//encoding the moc
	json.NewEncoder(w).Encode(movie)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteMovieById(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	/*remember think twice before using this function */
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
