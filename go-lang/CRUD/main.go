package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Modeil for course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

// fake db
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {
	fmt.Println("CRUD in go-lang")
	router := mux.NewRouter()

	// seeding
	courses = append(courses, Course{
		CourseName: "React",
		CourseId: "1",
		CoursePrice: 277,
		Author: &Author{
			Fullname: "Saurabh Dalakoti",
			Website: "dalakoti07.me",
		},
	})
	courses = append(courses, Course{
		CourseName: "Android",
		CourseId: "2",
		CoursePrice: 2772,
		Author: &Author{
			Fullname: "Saurabh Dalakoti",
			Website: "dalakoti07.me",
		},
	})

	// routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getAParticularCourse).Methods("GET")
	router.HandleFunc("/course", createCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")

	// start server
	log.Fatal(http.ListenAndServe(":4000",router))

}

// controllers - file

// serve home route
func serveHome(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte("<h1>Hello</h1>"))
}

func getAllCourses(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("get all course was called ... ")
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(courses)

}

func getAParticularCourse(writer http.ResponseWriter, r *http.Request) {
	fmt.Println("get a course was called ... ")
	writer.Header().Set("Content-Type", "application/json")

	// grab :id from request
	params := mux.Vars(r)

	// loop through courses
	for _, course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(writer).Encode(course)
			return 
		}
	}
	json.NewEncoder(writer).Encode("No course found with this id")
}


func createCourse(writer http.ResponseWriter, r *http.Request){
	fmt.Println("create course was called ... ")
	writer.Header().Set("Content-Type", "application/json")

	if r.Body == nil{
		json.NewEncoder(writer).Encode("Please send some data")
	}
	// {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(writer).Encode("No data in json")
		return
	}

	// generate id and append it
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses  = append(courses, course)
	json.NewEncoder(writer).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("create course was called ... ")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	
	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return 
		}
	}


}