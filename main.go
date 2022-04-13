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
	router.HandleFunc("/user", controllers.Authenticate(controllers.UpdateUser, 1)).Methods("PUT")
	router.HandleFunc("/user/hotel", controllers.GetHotelList).Methods("GET")
	router.HandleFunc("/user/hotel/room", controllers.GetRoomList).Methods("GET")
	// router.HandleFunc("/user/hotel/room", controllers.Authenticate(controllers.AddNewHotelOrder, 1)).Methods("POST")
	router.HandleFunc("/user/flight", controllers.GetFlightList).Methods("GET")
	router.HandleFunc("/user/flight/seat", controllers.GetFlightSeatList).Methods("GET")
	router.HandleFunc("/user/flight", controllers.Authenticate(controllers.AddNewFlightOrder, 1)).Methods("POST")
	// router.HandleFunc("/user/bus", controllers.GetBusList).Methods("GET")
	// router.HandleFunc("/user/bus", controllers.Authenticate(controllers.AddNewBusOrder, 1)).Methods("POST")
	// router.HandleFunc("/user/train", controllers.GetTrainList).Methods("GET")
	// router.HandleFunc("/user/train", controllers.Authenticate(controllers.AddNewTrainOrder, 1)).Methods("POST")
	router.HandleFunc("/user/tour", controllers.GetTourList).Methods("GET")
	router.HandleFunc("/user/tourschedule", controllers.GetTourScheduleList).Methods("GET")
	// router.HandleFunc("/user/tour", controllers.Authenticate(controllers.AddNewTourOrder, 1)).Methods("POST")
	// router.HandleFunc("/user/order", controllers.Authenticate(controllers.GetUserOrder, 1)).Methods("GET")
	// router.HandleFunc("/user/order", controllers.Authenticate(controllers.RequestRefund, 1)).Methods("PUT")

	router.HandleFunc("/partner", controllers.AddNewPartner).Methods("POST")
	router.HandleFunc("/partner", controllers.Authenticate(controllers.UpdatePartner, 2)).Methods("PUT")
	// router.HandleFunc("/partner/flight", controllers.Authenticate(controllers.GetFlightPartnerList, 2)).Methods("GET")
	// router.HandleFunc("/partner/flight", controllers.Authenticate(controllers.AddNewFlight, 2)).Methods("POST")
	// router.HandleFunc("/partner/flight", controllers.Authenticate(controllers.DeleteFlight, 2)).Methods("DELETE")

	router.HandleFunc("/admin/refund", controllers.Authenticate(controllers.GetRefundList, 0)).Methods("GET")
	// router.HandleFunc("/admin/refund", controllers.Authenticate(controllers.GetRefundList, 0)).Methods("GET")
	// router.HandleFunc("/admin/refund", controllers.Authenticate(controllers.ApproveRefund, 0)).Methods("DELETE")

	svrPort := controllers.LoadEnv("SVR_PORT")
	log.Println("Connected to port " + svrPort)
	addr := ":" + svrPort
	http.ListenAndServe(addr, router)
}
