package main

import (
	"fmt"
	"log"
	"net/http"
	c "uts_pbp/Controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("users", c.GetAllUsers).Methods("GET")
	// router.HandleFunc("users", c.InsertNewUser).Methods("POST")
	// router.HandleFunc("users", c.UpdateUser).Methods("PUT")
	// router.HandleFunc("users", c.Deleteuser).Methods("DELETE")

	http.Handle("v1/", router)
	fmt.Println("Connected to Port 8888")
	log.Println("Connected to Port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
