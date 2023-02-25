package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}

func getMovie(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(reader)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode("Not found")
}

func deleteMovie(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(reader)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index])
		}
	}
	json.NewEncoder(writer).Encode(movies)
}

func createMovie(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(reader.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func updateMovie(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(reader)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(reader.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(writer).Encode(movie)
			return
		}
	}
}

func main() {
	reader := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1", Isbn: "43121", Title: "Movie One", Director: &Director{
			FirstName: "Director 1 First Name",
			LastName:  "Director 1 last Name",
		},
	})
	movies = append(movies, Movie{
		ID: "2", Isbn: "43121", Title: "Movie Two", Director: &Director{
			FirstName: "Director 2 First Name",
			LastName:  "Director 2 Last Name",
		},
	})

	reader.HandleFunc("/movies", getMovies).Methods("GET")
	reader.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	reader.HandleFunc("/movies", createMovie).Methods("POST")
	reader.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	reader.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080")
	log.Fatal(
		http.ListenAndServe(":8080", reader))

}
