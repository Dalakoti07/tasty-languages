package main

import (
	"github.com/gorilla/mux"
	"go_book_management_system/pkg/routes"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRouters(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
