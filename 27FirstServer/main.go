package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for DB - file

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// fake DB

var courses = []Course{
	{
		CourseId:    "1",
		CourseName:  "DSA",
		CoursePrice: "499",
		Author:      &Author{Fullname: "Chandru", Website: "www.chandrudsa.com"},
	},
	{
		CourseId:    "2",
		CourseName:  "JS",
		CoursePrice: "299",
		Author:      &Author{Fullname: "Vicky", Website: "www.vickyjs.com"},
	},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", healthCheck).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}

// Controllers - file
func healthCheck(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Server is Running</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content/Type", "application/json")

	// data, _ := json.MarshalIndent(courses, "", "\t")
	// w.Write(data)

	json.NewEncoder(w).Encode(courses)

}
