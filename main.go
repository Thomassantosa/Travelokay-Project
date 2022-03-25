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

	svrPort := controllers.LoadEnv("SVR_PORT")
	log.Println("Connected to port " + svrPort)
	addr := ":" + svrPort
	http.ListenAndServe(addr, router)
}
