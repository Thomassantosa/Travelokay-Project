package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	models "github.com/Travelokay-Project/models"
	"github.com/joho/godotenv"
	// "github.com/Travelokay-Project/models"
)

func LoadEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Login(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	email := r.FormValue("Email")
	password := r.FormValue("Password")

	row := db.QueryRow("SELECT * FROM users WHERE email=? AND password=?", email, password)
	var user models.User
	if err := row.Scan(&user.ID, &user.Fullname, &user.Username, &user.Email, &user.Password, &user.Address, &user.UserType, &user.PartnerType, &user.CompanyName, &user.DateCreated); err != nil {
		log.Println(row)
		SendErrorResponse(w, 400)
		log.Print(err)
	} else {
		GenerateToken(w, user.ID, user.Username, user.UserType)
		SendSuccessResponse(w)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

	ResetUserToken(w)
	SendSuccessResponse(w)
}
func AddNewUser(w http.ResponseWriter, r *http.Request) {

	// connect to database
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	fullname := r.Form.Get("fullname")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	address := r.Form.Get("address")
	user_type := r.Form.Get("user_type")
	partner_type := r.Form.Get("partner_type")
	company_name := r.Form.Get("company_name")
	date_created := r.Form.Get("date_created")

	_, errQuery := db.Exec("INSERT INTO users(fullname,username,email,password,address,user_type,partner_type,company_name,date_created) values (?,?,?,?,?,?,?,?,?)", fullname, username, email, password, address, user_type, partner_type, company_name, date_created)

	if errQuery == nil {
		SendSuccessResponse(w)
	} else {
		SendErrorResponse(w, 400)
	}

	db.Close()
}
func UpdateUsers(w http.ResponseWriter, r *http.Request) {

	// connect to database
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	fullname := r.Form.Get("fullname")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	address := r.Form.Get("address")

	query := "UPDATE users SET"

	if fullname != "" {
		query += " fullname='" + fullname + "',"
	}
	if username != "" {
		query += " username='" + username + "',"
	}
	if password != "" {
		query += " password=" + password + ","
	}
	if address != "" {
		query += " address=" + address + ","
	}

	query1 := query[:len(query)-1]
	query1 += " WHERE email=" + email + "'"

	log.Println(query1)

	result, errQuery := db.Exec(query1)

	num, _ := result.RowsAffected()

	if errQuery == nil {
		if num == 0 {
			SendErrorResponse(w, 400)
		} else {
			SendSuccessResponse(w)
			log.Println(email)
		}
	} else {
		SendErrorResponse(w, 400)
	}

	db.Close()
}

func GetHotelList(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	hotelCity := r.Form.Get("hotel_city")

	rows, errQuery := db.Query("SELECT * FROM hotels WHERE hotel_city=?", hotelCity)

	var hotel models.Hotel
	var hotels []models.Hotel

	for rows.Next() {
		if err := rows.Scan(&hotel.ID, &hotel.HotelName, &hotel.HotelStar, &hotel.HotelRating, &hotel.HotelReview, &hotel.HotelFacility, &hotel.HotelAddress, &hotel.HotelCity, &hotel.HotelCountry); err != nil {
			log.Println(err.Error())
		} else {
			hotels = append(hotels, hotel)
		}
	}

	var response models.HotelsResponse
	if errQuery == nil {
		if len(hotels) == 0 {
			SendErrorResponse(w, 400)
		} else {
			response.Status = 200
			response.Message = "Success Get Data"
			response.Data = hotels
		}
	} else {
		SendErrorResponse(w, 400)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	db.Close()
}

func GetRoomList(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	hotelID, _ := strconv.Atoi(r.Form.Get("hotel_id"))

	rows, errQuery := db.Query("SELECT room_id,hotel_id,room_name,room_type,room_price,room_facility,room_capacity,room_status FROM rooms WHERE hotel_id=?", hotelID)

	var room models.Room
	var rooms []models.Room

	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.HotelID, &room.RoomName, &room.RoomType, &room.RoomPrice, &room.RoomFacility, &room.RoomCapacity, &room.RoomStatus); err != nil {
			log.Println(err.Error())
		} else {
			rooms = append(rooms, room)
		}
	}

	var response models.RoomsResponse
	if errQuery == nil {
		if len(rooms) == 0 {
			SendErrorResponse(w, 400)
		} else {
			response.Status = 200
			response.Message = "Success Get Data"
			response.Data = rooms
		}
	} else {
		SendErrorResponse(w, 400)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	db.Close()
}
func AddNewHotelOrder(w http.ResponseWriter, r *http.Request) {

	// connect to database
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	userID := r.Form.Get("user_id")
	roomID := r.Form.Get("room_id")
	orderStatus := r.Form.Get("order_status")
	orderDate := r.Form.Get("order_date")
	transactionType := r.Form.Get("transaction_type")

	_, errQuery := db.Exec("INSERT INTO orders(user_id,room_id,order_date,person_name,phone_number,order_status,transaction_type) values (?,?,?,?,?)", userID, roomID, orderDate, orderStatus, transactionType)

	if errQuery == nil {
		SendSuccessResponse(w)
	} else {
		SendErrorResponse(w, 400)
	}

	db.Close()
}

func GetFlightList(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	departureAirport, _ := strconv.Atoi(r.Form.Get("departure_airport"))

	rows, errQuery := db.Query("SELECT * FROM flights WHERE departure_airport=?", departureAirport)

	var flight models.Flight
	var flights []models.Flight

	for rows.Next() {
		if err := rows.Scan(&flight.ID, &flight.AirplaneID, &flight.DepartureAirport, &flight.DestinationAirport, &flight.FlightType, &flight.FlightNumber, &flight.DepartureTime, &flight.ArrivalTime, &flight.TravelTime); err != nil {
			log.Println(err.Error())
		} else {
			flights = append(flights, flight)
		}
	}

	var response models.FlightsResponse
	if errQuery == nil {
		if len(flights) == 0 {
			SendErrorResponse(w, 400)
		} else {
			response.Status = 200
			response.Message = "Success Get Data"
			response.Data = flights
		}
	} else {
		SendErrorResponse(w, 400)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	db.Close()
}

func GetBusList(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	departureBusstation, _ := strconv.Atoi(r.Form.Get("departure_busstation"))

	rows, errQuery := db.Query("SELECT * FROM bustrips WHERE departure_busstation=?", departureBusstation)

	var bus models.Bustrip
	var buses []models.Bustrip

	for rows.Next() {
		if err := rows.Scan(&bus.ID, &bus.BusID, &bus.DepartureBusstation, &bus.DestinationBusstation, &bus.BusNumber, &bus.DepartureTime, &bus.ArrivalTime, &bus.TravelTime); err != nil {
			log.Println(err.Error())
		} else {
			buses = append(buses, bus)
		}
	}

	var response models.BusesResponse
	if errQuery == nil {
		if len(buses) == 0 {
			SendErrorResponse(w, 400)
		} else {
			response.Status = 200
			response.Message = "Success Get Data"
			response.Data = buses
		}
	} else {
		SendErrorResponse(w, 400)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	db.Close()
}

func GetTrainList(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	departureStation, _ := strconv.Atoi(r.Form.Get("departure_station"))

	rows, errQuery := db.Query("SELECT * FROM traintrips WHERE departure_station=?", departureStation)

	var train models.Traintrip
	var trains []models.Traintrip

	for rows.Next() {
		if err := rows.Scan(&train.ID, &train.TrainID, &train.DepartureStation, &train.DestinationStation, &train.TraintripNumber, &train.DepartureTime, &train.ArrivalTime, &train.TravelTime); err != nil {
			log.Println(err.Error())
		} else {
			trains = append(trains, train)
		}
	}

	var response models.TrainsResponse
	if errQuery == nil {
		if len(trains) == 0 {
			SendErrorResponse(w, 400)
		} else {
			response.Status = 200
			response.Message = "Success Get Data"
			response.Data = trains
		}
	} else {
		SendErrorResponse(w, 400)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	db.Close()
}
