package main

import (
	"encoding/json"
	"fmt"
	"github.com/erezak/gocourse/c13rest/c13rest"
	"net/http"
	"log"
	"time"
)

func main() {

	// Create instance of rest client
	RestClient := c13rest.NewRestClient("https://gocourse.azure-mobile.net", "/tables/")

	dob, err := time.Parse("02/01/2006", "29/07/1972")
	if err != nil {
		log.Fatalf("Error while parsing date - %s", err)
		return
	}
	// Create a student
	student := c13rest.Student{StudentId: "029527474",
		Name:        "Erez A. Korn",
		DateOfBirth: dob}

	studentValue, err := json.Marshal(student)
	if err != nil {
		log.Fatalf("Error while encoding to json - %s", err)
		return
	}
	statusCode, status, response, err := RestClient.Create("students", studentValue)
	if err != nil {
		log.Fatalf("Error while creating student - %s", err)
		return
	}
	var studentResponse c13rest.StudentAzure
	err = json.Unmarshal(response, &studentResponse)
	if err != nil {
		log.Fatalf("Error while parsing response body - %s\nResponse:%s\n", err, response)
		return
	}

	var id = studentResponse.Id
	fmt.Printf("Create response - %d %s\n", statusCode, status)
	if statusCode == http.StatusCreated {
		fmt.Printf("Student created with Id %s\n", id)
	}

	// Delete the student
	statusCode, status, _, err = RestClient.Delete("students", id)
	if err != nil {
		log.Fatalf("Error while deleting student - %s", err)
		return
	}
	fmt.Printf("Delete response - %d %s\n", statusCode, status)

}
