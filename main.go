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
	router.HandleFunc("/user", controllers.AddNewUser).Methods("POST")
	router.HandleFunc("/user", controllers.UpdateUsers).Methods("PUT")
	router.HandleFunc("/user/hotel", controllers.GetHotelList).Methods("GET")
	router.HandleFunc("/user/hotel/room", controllers.GetRoomList).Methods("GET")
	router.HandleFunc("/user/hotel/room", controllers.AddNewHotelOrder).Methods("POST")
	router.HandleFunc("/user/flight", controllers.GetFlightList).Methods("GET")
	router.HandleFunc("/user/bus", controllers.GetBusList).Methods("GET")
	router.HandleFunc("/user/train", controllers.GetTrainList).Methods("GET")

	svrPort := controllers.LoadEnv("SVR_PORT")
	log.Println("Connected to port " + svrPort)
	addr := ":" + svrPort
	http.ListenAndServe(addr, router)
}
