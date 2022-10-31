package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/random"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`

}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item : = range movies{

		if item.ID == params["id"]{
		movies = append(movies[:index], movies[index+1:]...)
		break
	}

	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set(("content-Type","application/json"))
	params :=mux.vars(r)
	for _,item := range movies{
		if item.ID == params["id"]{
		json.newEncoder(w).Encoder(item)
	}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set(("content-Type","application/json"))
	
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	
}



func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID:"1",Isbn:"438227",Title:"Mvie One",Director :&Director{Firstname:"John", Lastname:"Doe"}})
	
	movies = append(movies, Movie{ID:"2",Isbn:"45455",Title:"Mvie Two",Director :&Director{Firstname:"Steve", Lastname:"Smith"}})
	r.HandleFunc("/movies", getMovies).Methods("Get")
	r.HandleFunc("/movies/{id}",getMovie).Methods("Get")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000,r"))

}