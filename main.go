package main

import (
	"log"
	"net/http"
	"os"

	controllers "github.com/Travelokay-Project/controllers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("RUNNING ...")

	router := mux.NewRouter()

	router.HandleFunc("/loginUser", controllers.LoginUser).Methods("POST")

	svrPort := os.Getenv("SVR_PORT")
	log.Println("Connected to port " + svrPort)
	addr := ":" + svrPort
	http.ListenAndServe(addr, router)

}
