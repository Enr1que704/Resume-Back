package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
}

type Job struct {
	Company     string         `json:"company"`
	Title       string         `json:"title"`
	Start       sql.NullString `json:"start"`
	End         sql.NullString `json:"end"`
	Description string         `json:"description"`
}

type Skill struct {
	Skill string `json:"skill"`
	Descr string `json:"descr"`
	Years int    `json:"years"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Home Page"}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	router.HandleFunc("/getJobs", jobHandler).Methods("GET")
	router.HandleFunc("/getSkills", skillHandler).Methods("GET")

	log.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

// TODO: Turn this into a datebase call at some point
func jobHandler(w http.ResponseWriter, r *http.Request) {
	// jobs := []Job{}
	// jobs = append(jobs, Job{Company: "Gideon Taylor Consulting", Title: "Intern", Start: "2022-01-10", End: "2024-05-10", Description: "Description 1"})
	// jobs = append(jobs, Job{Company: "Gideon Taylor Consulting", Title: "Software Engineer II", Start: "2024-05-10", End: "PRESENT", Description: "Description 2"})
	// json.NewEncoder(w).Encode(jobs)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	sqlPath := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/resume", dbUser, dbPass)

	db, dbErr := sql.Open("mysql", sqlPath)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM job")

	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var job Job
		err = results.Scan(&job.Company, &job.Title, &job.Start, &job.End, &job.Description)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(job)
	}
}

func skillHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	sqlPath := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/resume", dbUser, dbPass)

	db, dbErr := sql.Open("mysql", sqlPath)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM skill")

	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var skill Skill
		err = results.Scan(&skill.Skill, &skill.Descr, &skill.Years)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(skill)
	}
}
