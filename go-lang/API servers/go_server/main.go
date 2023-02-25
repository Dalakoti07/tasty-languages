package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(writer http.ResponseWriter, reader *http.Request) {
	err := reader.ParseForm()
	if err != nil {
		fmt.Fprintf(writer, "parseForm Error: %v", err)
		return
	}
	fmt.Fprintf(writer, "Post request successful")
	name := reader.FormValue("name")
	address := reader.FormValue("address")
	fmt.Fprintf(writer, "Name= %s\n", name)
	fmt.Fprintf(writer, "Address= %s\n", address)
}

func helloHandler(writer http.ResponseWriter, reader *http.Request) {
	if reader.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}
	if reader.Method != "GET" {
		http.Error(writer, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(writer, "hello! from golang")
}

/**
Servers 3 endpoints
- '/' which will serve static website
- '/form' which servers html form page
- '/hello' which would return hello nothing else in response
*/

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
