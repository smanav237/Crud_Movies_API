package main

import(
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json: "id"`               // json for testing the api
	Title string `json: "title"`
	Actor string `json: "actor`
}
var movies []Movie    // slice of movies

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")   // sets the "Content-Type" header of the HTTP response to "application/json
	params := mux.Vars(r)             //  to extract route parameters from the request
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)     // delete the curr movie and append rest of movies in array
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {      // _ bcoz no index
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	// uses json.NewDecoder to decode the JSON data from the request body (r.Body) into the movie variable.
	_ = json.NewDecoder(r.Body).Decode(&movie)    // _ is used to discard the error returned by 'Decode'
	
	movies = append(movies, movie)
	
	json.NewEncoder(w).Encode(movie)    // encodes the created movie into JSON format and writes it to the response body.
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// here we are first deleting the curr movie (with id) and then adding the new movie(which is sent in body of postman)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)   // deletion
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)      // similar to creation
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func main(){
	r := mux.NewRouter()    // initialize new router using gorilla mux package
	
	movies = append(movies, Movie{ID: "1", Title: "movie1", Actor: "Leonardo Dicaprio"})
	movies = append(movies, Movie{ID: "2", Title: "movie2", Actor: "Christian Bale"})
	movies = append(movies, Movie{ID: "3", Title: "movie3", Actor: "Ryan Gosling"})
	movies = append(movies, Movie{ID: "4", Title: "movie4", Actor: "Denzel Washington"})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server successfully started at PORT: 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))

}