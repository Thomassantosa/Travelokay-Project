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

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
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

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	fullname := r.Form.Get("fullname")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	address := r.Form.Get("address")
	partnerType := r.Form.Get("partnerType")
	companyName := r.Form.Get("companyName")

	// Encrypt password
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	// Query
	_, errQuery := db.Exec("INSERT INTO users(fullname, username, email, password, address, user_type, partner_type, company_name) values (?,?,?,?,?,2,?,?)", fullname, username, email, encryptedPassword, address, partnerType, companyName)

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
		if err := rows.Scan(&hotel.ID, &hotel.HotelName, &hotel.HotelStar, &hotel.HotelRating, &hotel.HotelReview, &hotel.HotelAddress, &hotel.HotelFacility, &hotel.HotelCity, &hotel.HotelCountry); err != nil {
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

// func AddNewHotelOrder(w http.ResponseWriter, r *http.Request) {

// 	// connect to database
// 	db := Connect()
// 	defer db.Close()

// 	err := r.ParseForm()
// 	if err != nil {
// 		return
// 	}
// 	userID := r.Form.Get("user_id")
// 	roomID := r.Form.Get("room_id")
// 	email := r.Form.Get("email")
// 	phoneNumber := r.Form.Get("phone_number")
// 	transactionType := r.Form.Get("transaction_type")
// 	personName := r.Form.Get("person_name")
// 	orderDate := r.Form.Get("order_date")

// 	_, errQuery := db.Exec("INSERT INTO orders(user_id,room_id,order_date,person_name,phone_number,email,transaction_type) values (?,?,?,?,?,?,?)", userID, roomID, orderDate, personName, phoneNumber, email, transactionType)

// 	if errQuery == nil {
// 		SendSuccessResponse(w)
// 	} else {
// 		SendErrorResponse(w, 400)
// 	}

// 	db.Close()
// }

func GetFlightList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from query params
	departureCity := r.URL.Query().Get("departureCity")
	destinationCity := r.URL.Query().Get("destinationCity")
	seatType := r.URL.Query().Get("seatType")
	departureDate := r.URL.Query().Get("departureDate")

	query :=
		`SELECT flights.flight_id, airplanes.airplane_model, airlines.airline_name, airportA.airport_id, airportA.airport_code,` +
			` airportA.airport_name, airportA.airport_city, airportA.airport_country, airportB.airport_id, airportB.airport_code,` +
			` airportB.airport_name, airportB.airport_city, airportB.airport_country, flight_type, flight_number, departure_time,` +
			` arrival_time, travel_time FROM flights` +
			` JOIN airplanes ON flights.airplane_id = airplanes.airplane_id` +
			` JOIN airlines ON airplanes.airline_id = airlines.airline_id` +
			` JOIN airports AS airportA ON flights.departure_airport = airportA.airport_id` +
			` JOIN airports AS airportB ON flights.destination_airport = airportB.airport_id` +
			` JOIN seats ON flights.flight_id = seats.flight_id` +
			` WHERE airportA.airport_city = ? AND` +
			` airportB.airport_city = ? AND` +
			` CAST(departure_time AS DATE) = ? AND` +
			` seats.seat_type = ?` +
			` GROUP BY flights.flight_id`

	// rows, errQuery := db.Query(query, departureCity, destinationCity, departureDate, seatType)
	rows, errQuery := db.Query(query, departureCity, destinationCity, departureDate, seatType)

	if errQuery != nil {
		SendErrorResponse(w, 500)
		log.Println(errQuery)
		return
	}

	var flight models.Flight
	var flights []models.Flight

	for rows.Next() {
		err := rows.Scan(&flight.ID, &flight.AirplaneModel, &flight.AirlineName, &flight.DepartureAirport.ID, &flight.DepartureAirport.Code,
			&flight.DepartureAirport.Name, &flight.DepartureAirport.City, &flight.DepartureAirport.Country, &flight.DestinationAirport.ID,
			&flight.DestinationAirport.Code, &flight.DestinationAirport.Name, &flight.DestinationAirport.City,
			&flight.DestinationAirport.Country, &flight.FlightType, &flight.FlightNumber, &flight.DepartureTime, &flight.ArrivalTime,
			&flight.TravelTime)
		if err != nil {
			SendErrorResponse(w, 500)
			log.Println(err)
			return
		} else {
			flights = append(flights, flight)
		}
	}

	var response models.FlightsResponse
	if errQuery == nil {
		if len(flights) == 0 {
			SendErrorResponse(w, 400)
			return
		} else {
			response.Status = 200
			response.Message = "Get Data Success"
			response.Data = flights
		}
	} else {
		SendErrorResponse(w, 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetFlightSeatList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from query params
	flightId := r.URL.Query().Get("flightId")
	seatType := r.URL.Query().Get("seatType")

	query :=
		`SELECT seat_id, seat_type, seat_name, seat_status, baggage_capacity, seat_price FROM seats` +
			` WHERE flight_id = ? AND` +
			` seat_type = ? AND` +
			` seat_status = 0`

	rows, errQuery := db.Query(query, flightId, seatType)

	if errQuery != nil {
		SendErrorResponse(w, 500)
		log.Println(errQuery)
		return
	}

	var seat models.Seat
	var seats []models.Seat

	for rows.Next() {
		err := rows.Scan(&seat.ID, &seat.SeatType, &seat.SeatName, &seat.SeatStatus, &seat.BaggageCapacity, &seat.SeatPrice)
		if err != nil {
			SendErrorResponse(w, 500)
			log.Println(err)
			return
		} else {
			seats = append(seats, seat)
		}
	}

	var response models.SeatsResponse
	if errQuery == nil {
		if len(seats) == 0 {
			SendErrorResponse(w, 400)
			return
		} else {
			response.Status = 200
			response.Message = "Get Data Success"
			response.Data = seats
		}
	} else {
		SendErrorResponse(w, 400)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// func GetBusList(w http.ResponseWriter, r *http.Request) {
// 	db := Connect()
// 	defer db.Close()

// 	err := r.ParseForm()
// 	if err != nil {
// 		return
// 	}

// 	departureBusstation, _ := strconv.Atoi(r.Form.Get("departure_busstation"))

// 	rows, errQuery := db.Query("SELECT * FROM bustrips WHERE departure_busstation=?", departureBusstation)

// 	var bus models.Bustrip
// 	var buses []models.Bustrip

// 	for rows.Next() {
// 		if err := rows.Scan(&bus.ID, &bus.BusID, &bus.DepartureBusstation, &bus.DestinationBusstation, &bus.BusNumber, &bus.DepartureTime, &bus.ArrivalTime, &bus.DepartureDate, &bus.ArrivalDate, &bus.TravelTime); err != nil {
// 			log.Println(err.Error())
// 		} else {
// 			buses = append(buses, bus)
// 		}
// 	}

// 	var response models.BusesResponse
// 	if errQuery == nil {
// 		if len(buses) == 0 {
// 			SendErrorResponse(w, 400)
// 		} else {
// 			response.Status = 200
// 			response.Message = "Success Get Data"
// 			response.Data = buses
// 		}
// 	} else {
// 		SendErrorResponse(w, 400)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// 	db.Close()
// }

// func GetTrainList(w http.ResponseWriter, r *http.Request) {
// 	db := Connect()
// 	defer db.Close()

// 	err := r.ParseForm()
// 	if err != nil {
// 		return
// 	}

// 	departureStation, _ := strconv.Atoi(r.Form.Get("departure_station"))

// 	rows, errQuery := db.Query("SELECT * FROM traintrips WHERE departure_station=?", departureStation)

// 	var train models.Traintrip
// 	var trains []models.Traintrip

// 	for rows.Next() {
// 		if err := rows.Scan(&train.ID, &train.TrainID, &train.DepartureStation, &train.DestinationStation, &train.TraintripNumber, &train.DepartureTime, &train.ArrivalTime, &train.DepartureDate, &train.ArrivalDate, &train.TravelTime); err != nil {
// 			log.Println(err.Error())
// 		} else {
// 			trains = append(trains, train)
// 		}
// 	}

// 	var response models.TrainsResponse
// 	if errQuery == nil {
// 		if len(trains) == 0 {
// 			SendErrorResponse(w, 400)
// 		} else {
// 			response.Status = 200
// 			response.Message = "Success Get Data"
// 			response.Data = trains
// 		}
// 	} else {
// 		SendErrorResponse(w, 400)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// 	db.Close()
// }

func GetTourList(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := Connect()
	defer db.Close()

	tourCity := r.URL.Query().Get("tourCity")

	rows, errQuery := db.Query("SELECT * FROM tours WHERE tour_city=?", tourCity)
	if errQuery != nil {
		SendErrorResponse(w, 400)
		return
	}
	var tour models.Tours
	var tours []models.Tours

	for rows.Next() {
		if err := rows.Scan(&tour.ID, &tour.TourName, &tour.TourRating, &tour.TourReview, &tour.TourDesc, &tour.TourFacility, &tour.TourAddress, &tour.TourCity, &tour.TourProvince, &tour.TourCountry); err != nil {
			log.Println(err.Error())
			return
		} else {
			tours = append(tours, tour)
		}
	}

	var response models.ToursResponse

	if len(tours) == 0 {
		SendErrorResponse(w, 204)
		return
	} else {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = tours
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func GetTourScheduleList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	tourId := r.URL.Query().Get("tourId")

	rows, errQuery := db.Query("SELECT * FROM tourschedules WHERE tour_id=?", tourId)
	if errQuery != nil {
		SendErrorResponse(w, 400)
		return
	}
	var tour models.ToursSchedule
	var tours []models.ToursSchedule

	for rows.Next() {
		if err := rows.Scan(&tour.ID, &tour.TourID, &tour.ScheduleDay, &tour.OpenTime, &tour.CloseTime, &tour.Price); err != nil {
			log.Println(err.Error())
			return
		} else {
			tours = append(tours, tour)
		}
	}

	var response models.ToursScheduleResponse
	if len(tours) == 0 {
		SendErrorResponse(w, 204)
		return
	} else {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = tours
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}
