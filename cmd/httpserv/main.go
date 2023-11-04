package main

import (
	"github.com/filatkinen/cassandra/internal"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	workers = 1
	records = 10
)

func main() {
	defer internal.Session.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", internal.HomeLink)
	router.HandleFunc("/create", internal.CreateStudent).Methods("POST")              // http://localhost:3000/create
	router.HandleFunc("/getstudents", internal.GetAllStudents).Methods("GET")         // http://localhost:3000/getstudents
	router.HandleFunc("/count", internal.CountAllStudents).Methods("GET")             // http://localhost:3000/count
	router.HandleFunc("/getone/{id}", internal.GetOneStudent).Methods("GET")          // http://localhost:3000/getone/1
	router.HandleFunc("/deleteone/{id}", internal.DeleteOneStudent).Methods("DELETE") // http://localhost:3000/deleteone/1
	router.HandleFunc("/deleteall", internal.DeleteAllStudents).Methods("DELETE")     // http://localhost:3000/deleteall
	router.HandleFunc("/update/{id}", internal.UpdateStudent).Methods("PATCH")        // http://localhost:3000/update/3
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))

}
