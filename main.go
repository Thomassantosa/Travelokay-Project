package main

import (
	"log"
	"net/http"

	controllers "github.com/Travelokay-Project/controllers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("RUNNING ...")

	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")
	router.HandleFunc("/user", controllers.InsertUsers).Methods("POST")
	router.HandleFunc("/user", controllers.UpdateUsers).Methods("PUT")

	svrPort := controllers.LoadEnv("SVR_PORT")
	log.Println("Connected to port " + svrPort)
	addr := ":" + svrPort
	http.ListenAndServe(addr, router)
}
