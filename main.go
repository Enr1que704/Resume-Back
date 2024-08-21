package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

type Job struct {
	Company     string `json:"company"`
	Title       string `json:"title"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Description string `json:"description"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Home Page"}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	router.HandleFunc("/getJobs", jobHandler).Methods("GET")

	log.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// TODO: Turn this into a datebase call at some point
func jobHandler(w http.ResponseWriter, r *http.Request) {
	jobs := []Job{}
	jobs = append(jobs, Job{Company: "Gideon Taylor Consulting", Title: "Intern", Start: "2022-01-10", End: "2024-05-10", Description: "Description 1"})
	jobs = append(jobs, Job{Company: "Gideon Taylor Consulting", Title: "Software Engineer II", Start: "2024-05-10", End: "PRESENT", Description: "Description 2"})

	json.NewEncoder(w).Encode(jobs)
}
