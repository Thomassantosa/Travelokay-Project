package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	models "github.com/Travelokay-Project/models"
	"github.com/joho/godotenv"
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

	// Get value from form
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Encrypt password
	// hasher := md5.New()
	// hasher.Write([]byte(password))
	// encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	// Query
	row := db.QueryRow("SELECT user_type FROM users WHERE email=? AND password=?", email, password)
	var userType int
	if err := row.Scan(&userType); err != nil {
		SendErrorResponse(w, 400)
		log.Print(err)
		log.Print("(ERROR) email or username not found")
	} else {

		if userType == 2 {
			row := db.QueryRow("SELECT * FROM users WHERE email=? AND password=?", email, password)
			var partner models.Partner
			if err := row.Scan(&partner.ID, &partner.Fullname, &partner.Username, &partner.Email, &partner.Password, &partner.Address, &partner.UserType, &partner.PartnerType, &partner.CompanyName, &partner.DateCreated); err != nil {
				SendErrorResponse(w, 400)
				log.Print("A")
				log.Print(err)
			} else {
				GenerateToken(w, partner.ID, partner.Username, partner.UserType)

				// Response
				var partnerResponse models.PartnerResponse
				partnerResponse.Status = 200
				partnerResponse.Message = "Request success"
				partnerResponse.Data = partner

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(partnerResponse)
			}
		} else {
			row := db.QueryRow("SELECT user_id, fullname, username, email, password, address, user_type, date_created FROM users WHERE email=? AND password=?", email, password)
			var user models.User
			if err := row.Scan(&user.ID, &user.Fullname, &user.Username, &user.Email, &user.Password, &user.Address, &user.UserType, &user.DateCreated); err != nil {
				SendErrorResponse(w, 400)
				log.Print(err)
				log.Print("B")
			} else {
				GenerateToken(w, user.ID, user.Username, user.UserType)

				// Response
				var userResponse models.UserResponse
				userResponse.Status = 200
				userResponse.Message = "Request success"
				userResponse.Data = user

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(userResponse)
			}

		}
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

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		return
	}
	fullname := r.Form.Get("fullname")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	address := r.Form.Get("address")

	// Encrypt password
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	// Query
	_, errQuery := db.Exec("INSERT INTO users(fullname, username, email, password, address, user_type) values (?,?,?,?,?,1)", fullname, username, email, encryptedPassword, address)

	if errQuery == nil {
		SendSuccessResponse(w)
	} else {
		SendErrorResponse(w, 400)
	}
}

func AddNewPartner(w http.ResponseWriter, r *http.Request) {

	// connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		return
	}
	fullname := r.Form.Get("fullname")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	address := r.Form.Get("address")
	partner_type := r.Form.Get("partner_type")
	company_name := r.Form.Get("company_name")

	// Encrypt password
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	// Query
	_, errQuery := db.Exec("INSERT INTO users(fullname, username, email, password, address, user_type, partner_type, company_name) values (?,?,?,?,?,2,?,?)", fullname, username, email, encryptedPassword, address, partner_type, company_name)

	if errQuery == nil {
		SendSuccessResponse(w)
	} else {
		SendErrorResponse(w, 400)
	}
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
		if err := rows.Scan(&hotel.ID, &hotel.HotelName, &hotel.HotelStar, &hotel.HotelReview, &hotel.HotelRating, &hotel.HotelAddress, &hotel.HotelFacility, &hotel.HotelCity, &hotel.HotelCountry); err != nil {
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
	email := r.Form.Get("email")
	phoneNumber := r.Form.Get("phone_number")
	transactionType := r.Form.Get("transaction_type")
	personName := r.Form.Get("person_name")
	orderDate := r.Form.Get("order_date")

	_, errQuery := db.Exec("INSERT INTO orders(user_id,room_id,order_date,person_name,phone_number,email,transaction_type) values (?,?,?,?,?,?,?)", userID, roomID, orderDate, personName, phoneNumber, email, transactionType)

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
		if err := rows.Scan(&flight.ID, &flight.AirplaneID, &flight.DepartureAirport, &flight.DestinationAirport, &flight.FlightType, &flight.FlightNumber, &flight.DepartureTime, &flight.ArrivalTime, &flight.DepartureDate, &flight.ArrivalDate, &flight.TravelTime); err != nil {
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
		if err := rows.Scan(&bus.ID, &bus.BusID, &bus.DepartureBusstation, &bus.DestinationBusstation, &bus.BusNumber, &bus.DepartureTime, &bus.ArrivalTime, &bus.DepartureDate, &bus.ArrivalDate, &bus.TravelTime); err != nil {
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
		if err := rows.Scan(&train.ID, &train.TrainID, &train.DepartureStation, &train.DestinationStation, &train.TraintripNumber, &train.DepartureTime, &train.ArrivalTime, &train.DepartureDate, &train.ArrivalDate, &train.TravelTime); err != nil {
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
