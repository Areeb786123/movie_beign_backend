package controllers

import (
	"context"
	"corses/entity"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://areeb_user_31:786Areeb@moviesbeign.o2b1e.mongodb.net/"
const dbName = "moviesVerse"
const collectionList = "movieList"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(collectionList)
	fmt.Println("collection instance is ready")
}

func getMovieById(movieId string) {

}
func getAllMovies() {

}

func updateMovie(movieId string) {

}

func deleteMovieById(movieId string) {

}

func deleteAllMovies() {

}

func addMovies(movies entity.MovieEntity) {

}

func AddMovie(w http.ResponseWriter, r *http.Request) {

}

func GetMoviesById(w http.ResponseWriter, r *http.Request) {

}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {

}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	/*remember think twice before using this function */
}
