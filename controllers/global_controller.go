package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"

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
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	// Query
	row := db.QueryRow("SELECT user_type FROM users WHERE email=? AND password=?", email, encryptedPassword)
	var userType int
	if err := row.Scan(&userType); err != nil {
		log.Println("(ERROR)\t", err)
		SendErrorResponse(w, 400)
		return
	}

	if userType == 2 {

		// Partner login
		row := db.QueryRow("SELECT * FROM users WHERE email=? AND password=?", email, encryptedPassword)
		var partner models.Partner
		if err := row.Scan(&partner.ID, &partner.Fullname, &partner.Username, &partner.Email, &partner.Password, &partner.Address, &partner.UserType, &partner.PartnerType, &partner.CompanyName, &partner.DateCreated); err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 400)
			return
		}

		GenerateToken(w, partner.ID, partner.Username, partner.UserType)

		// Response
		var partnerResponse models.PartnerResponse
		partnerResponse.Status = 200
		partnerResponse.Message = "Request success"
		partnerResponse.Data = partner

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(partnerResponse)

	} else {

		// User / admin login
		row := db.QueryRow("SELECT user_id, fullname, username, email, password, address, user_type, date_created FROM users WHERE email=? AND password=?", email, encryptedPassword)
		var user models.User
		if err := row.Scan(&user.ID, &user.Fullname, &user.Username, &user.Email, &user.Password, &user.Address, &user.UserType, &user.DateCreated); err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 400)
			return
		}

		GenerateToken(w, user.ID, user.Username, user.UserType)

		// Response
		var userResponse models.UserResponse
		userResponse.Status = 200
		userResponse.Message = "Request success"
		userResponse.Data = user

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userResponse)
	}
	log.Println("(SUCCESS)\t", "Login request")
}

func Logout(w http.ResponseWriter, r *http.Request) {

	ResetUserToken(w)
	SendSuccessResponse(w)
	log.Println("(SUCCESS)\t", "Logout request")
}

func AddNewUser(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	fullname := r.FormValue("fullname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	address := r.FormValue("address")

	// Encrypt password
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	// Query
	_, errQuery := db.Exec("INSERT INTO users(fullname, username, email, password, address, user_type) values (?,?,?,?,?,1)", fullname, username, email, encryptedPassword, address)

	if errQuery == nil {
		log.Println("(SUCCESS)\t", "Add new user request")
		SendSuccessResponse(w)
	} else {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
	}
}

func AddNewPartner(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	fullname := r.FormValue("fullname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	address := r.FormValue("address")
	partnerType := r.FormValue("partnerType")
	companyName := r.FormValue("companyName")

	// Encrypt password
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	// Query
	_, errQuery := db.Exec("INSERT INTO users(fullname, username, email, password, address, user_type, partner_type, company_name) values (?,?,?,?,?,2,?,?)", fullname, username, email, encryptedPassword, address, partnerType, companyName)

	if errQuery == nil {
		log.Println("(SUCCESS)\t", "Add new partner request")
		SendSuccessResponse(w)
	} else {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
	}
}

func GetHotelList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from query params
	hotelCity := r.URL.Query().Get("hotel_city")

	// Query
	rows, errQuery := db.Query("SELECT * FROM hotels WHERE hotel_city = ?", hotelCity)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 500)
		return
	}

	// Set value
	var hotel models.Hotel
	var hotels []models.Hotel

	for rows.Next() {
		if err := rows.Scan(&hotel.ID, &hotel.HotelName, &hotel.HotelStar, &hotel.HotelRating, &hotel.HotelReview, &hotel.HotelAddress, &hotel.HotelFacility, &hotel.HotelCity, &hotel.HotelCountry); err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 500)
			return
		}
		hotels = append(hotels, hotel)
	}

	// Response
	var response models.HotelsResponse

	if len(hotels) == 0 {
		log.Println("(ERROR)\t", "Data empty")
		SendErrorResponse(w, 400)
	} else {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = hotels

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		log.Println("(SUCCESS)\t", "Get hotel list request")
	}
}

func GetRoomList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from query params
	hotelID := r.URL.Query().Get("hotel_id")

	// Query
	rows, errQuery := db.Query("SELECT room_id, hotel_id, room_name, room_type, room_price, room_facility, room_capacity, room_status FROM rooms WHERE hotel_id = ?", hotelID)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 500)
		return
	}

	// Set value
	var room models.Room
	var rooms []models.Room

	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.HotelID, &room.RoomName, &room.RoomType, &room.RoomPrice, &room.RoomFacility, &room.RoomCapacity, &room.RoomStatus); err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 500)
			return
		}
		rooms = append(rooms, room)
	}

	// Response
	var response models.RoomsResponse

	if len(rooms) == 0 {
		log.Println("(ERROR)\t", "Data empty")
		SendErrorResponse(w, 400)
	} else {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = rooms

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		log.Println("(SUCCESS)\t", "Get room list request")
	}
}

func GetFlightList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from query params
	departureCity := r.URL.Query().Get("departureCity")
	destinationCity := r.URL.Query().Get("destinationCity")
	seatType := r.URL.Query().Get("seatType")
	departureDate := r.URL.Query().Get("departureDate")
	log.Println(departureCity)
	log.Println(destinationCity)
	log.Println(seatType)
	log.Println(departureDate)

	// Query
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

	rows, errQuery := db.Query(query, departureCity, destinationCity, departureDate, seatType)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 500)
		return
	}

	// Set values
	var flight models.Flight
	var flights []models.Flight

	for rows.Next() {
		err := rows.Scan(&flight.ID, &flight.AirplaneModel, &flight.AirlineName, &flight.DepartureAirport.ID, &flight.DepartureAirport.Code,
			&flight.DepartureAirport.Name, &flight.DepartureAirport.City, &flight.DepartureAirport.Country, &flight.DestinationAirport.ID,
			&flight.DestinationAirport.Code, &flight.DestinationAirport.Name, &flight.DestinationAirport.City,
			&flight.DestinationAirport.Country, &flight.FlightType, &flight.FlightNumber, &flight.DepartureTime, &flight.ArrivalTime,
			&flight.TravelTime)

		if err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 500)
			return
		}
		flights = append(flights, flight)
	}

	// Response
	var response models.FlightsResponse

	if len(flights) == 0 {
		log.Println("(ERROR)\t", "Data empty")
		SendErrorResponse(w, 204)
	} else {
		response.Status = 200
		response.Message = "Get Data Success"
		response.Data = flights

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		log.Println("(SUCCESS)\t", "Get flight list request")
	}
}

func GetSeatList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from query params
	flightId := r.URL.Query().Get("flightId")
	trainTripId := r.URL.Query().Get("trainTripId")
	BusTripId := r.URL.Query().Get("busTripId")
	seatType := r.URL.Query().Get("seatType")

	// Query
	query :=
		`SELECT seat_id, seat_type, seat_name, seat_status, baggage_capacity, seat_price FROM seats` +
			` WHERE seat_status = 0 AND` +
			` seat_type = ? AND`

	if flightId != "" {
		query += " flight_id = " + flightId
	}
	if trainTripId != "" {
		query += " traintrip_id = " + trainTripId
	}
	if BusTripId != "" {
		query += " bustrip_id = " + BusTripId
	}

	rows, errQuery := db.Query(query, seatType)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 500)
		return
	}

	// Set values
	var seat models.Seat
	var seats []models.Seat

	for rows.Next() {
		err := rows.Scan(&seat.ID, &seat.SeatType, &seat.SeatName, &seat.SeatStatus, &seat.BaggageCapacity, &seat.SeatPrice)
		if err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 500)
			return
		}
		seats = append(seats, seat)
	}

	// Response
	var response models.SeatsResponse

	if len(seats) == 0 {
		log.Println("(ERROR)\t", "Data empty")
		SendErrorResponse(w, 204)
	} else {
		response.Status = 200
		response.Message = "Get Data Success"
		response.Data = seats

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		log.Println("(SUCCESS)\t", "Get seat list request")
	}
}

func GetBusTripList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from query params
	departureCity := r.URL.Query().Get("departureCity")
	destinationCity := r.URL.Query().Get("destinationCity")
	seatType := r.URL.Query().Get("seatType")
	departureDate := r.URL.Query().Get("departureDate")

	// Query
	query :=
		`SELECT bustrips.bustrip_id, buses.bus_model, buscompanies.buscompany_name, busstationA.busstation_id, busstationA.busstation_code,` +
			` busstationA.busstation_name, busstationA.busstation_city, busstationB.busstation_id, busstationB.busstation_code,` +
			` busstationB.busstation_name, busstationB.busstation_city, bustrip_number, departure_time, arrival_time, travel_time FROM bustrips` +
			` JOIN buses ON bustrips.bus_id = buses.bus_id` +
			` JOIN buscompanies ON buses.buscompany_id = buscompanies.buscompany_id` +
			` JOIN busstations AS busstationA ON bustrips.departure_busstation = busstationA.busstation_id` +
			` JOIN busstations AS busstationB ON bustrips.destination_busstation = busstationB.busstation_id` +
			` JOIN seats ON bustrips.bustrip_id = seats.bustrip_id` +
			` WHERE busstationA.busstation_city = ? AND` +
			` busstationB.busstation_city = ? AND` +
			` CAST(departure_time AS DATE) = ? AND` +
			` seats.seat_type = ?` +
			` GROUP BY bustrips.bustrip_id`

	rows, errQuery := db.Query(query, departureCity, destinationCity, departureDate, seatType)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 500)
		return
	}

	// Set values
	var busTrip models.Bustrip
	var busTrips []models.Bustrip

	for rows.Next() {
		err := rows.Scan(&busTrip.ID, &busTrip.BusModel, &busTrip.CompanyName, &busTrip.DepartureBusstation.ID, &busTrip.DepartureBusstation.Code,
			&busTrip.DepartureBusstation.Name, &busTrip.DepartureBusstation.City, &busTrip.DestinationBusstation.ID,
			&busTrip.DestinationBusstation.Code, &busTrip.DestinationBusstation.Name, &busTrip.DestinationBusstation.City, &busTrip.BustripNumber,
			&busTrip.DepartureTime, &busTrip.ArrivalTime, &busTrip.TravelTime)

		if err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 500)
			return
		}
		busTrips = append(busTrips, busTrip)
	}

	// Response
	var response models.BusesResponse

	if len(busTrips) == 0 {
		log.Println("(ERROR)\t", "Data empty")
		SendErrorResponse(w, 204)
	} else {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = busTrips

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		log.Println("(SUCCESS)\t", "Get bustrip list request")
	}
}

func GetTrainTripList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from query params
	departureCity := r.URL.Query().Get("departureCity")
	destinationCity := r.URL.Query().Get("destinationCity")
	seatType := r.URL.Query().Get("seatType")
	departureDate := r.URL.Query().Get("departureDate")

	// Query
	query :=
		`SELECT traintrips.traintrip_id, trains.train_model, stationA.station_id, stationA.station_code,` +
			` stationA.station_name, stationA.station_city, stationB.station_id, stationB.station_code,` +
			` stationB.station_name, stationB.station_city, traintrip_number, departure_time, arrival_time, travel_time FROM traintrips` +
			` JOIN trains ON traintrips.train_id = trains.train_id` +
			` JOIN stations AS stationA ON traintrips.departure_station = stationA.station_id` +
			` JOIN stations AS stationB ON traintrips.destination_station = stationB.station_id` +
			` JOIN seats ON traintrips.traintrip_id = seats.traintrip_id` +
			` WHERE stationA.station_city = ? AND` +
			` stationB.station_city = ? AND` +
			` CAST(departure_time AS DATE) = ? AND` +
			` seats.seat_type = ?` +
			` GROUP BY traintrips.traintrip_id`

	rows, errQuery := db.Query(query, departureCity, destinationCity, departureDate, seatType)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 500)
		return
	}

	// Set values
	var trainTrip models.Traintrip
	var trainTrips []models.Traintrip

	for rows.Next() {
		err := rows.Scan(&trainTrip.ID, &trainTrip.TrainModel, &trainTrip.DepartureStation.ID, &trainTrip.DepartureStation.Code,
			&trainTrip.DepartureStation.Name, &trainTrip.DepartureStation.City, &trainTrip.DestinationStation.ID,
			&trainTrip.DestinationStation.Code, &trainTrip.DestinationStation.Name, &trainTrip.DestinationStation.City, &trainTrip.TrainTripNumber,
			&trainTrip.DepartureTime, &trainTrip.ArrivalTime, &trainTrip.TravelTime)

		if err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 500)
			return
		}
		trainTrips = append(trainTrips, trainTrip)
	}

	// Response
	var response models.TrainTripsResponse

	if len(trainTrips) == 0 {
		log.Println("(ERROR)\t", "Data empty")
		SendErrorResponse(w, 204)
	} else {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = trainTrips

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		log.Println("(SUCCESS)\t", "Get traitTrip list request")
	}
}

func GetTourList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from query params
	tourCity := r.URL.Query().Get("tourCity")

	// Query
	rows, errQuery := db.Query("SELECT * FROM tours WHERE tour_city = ?", tourCity)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
		return
	}

	// Set value
	var tour models.Tours
	var tours []models.Tours

	for rows.Next() {
		if err := rows.Scan(&tour.ID, &tour.TourName, &tour.TourRating, &tour.TourReview, &tour.TourDesc, &tour.TourFacility, &tour.TourAddress, &tour.TourCity, &tour.TourProvince, &tour.TourCountry); err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 400)
			return
		}
		tours = append(tours, tour)
	}

	var response models.ToursResponse

	if len(tours) == 0 {
		log.Println("(ERROR)\t", "Data empty")
		SendErrorResponse(w, 204)
		return
	} else {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = tours

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		log.Println("(SUCCESS)\t", "Get tour list request")
	}
}

func GetTourScheduleList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from query params
	tourId := r.URL.Query().Get("tourId")

	// Query
	rows, errQuery := db.Query("SELECT * FROM tourschedules WHERE tour_id = ?", tourId)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
		return
	}

	// Set value
	var tour models.ToursSchedule
	var tours []models.ToursSchedule

	for rows.Next() {
		if err := rows.Scan(&tour.ID, &tour.TourID, &tour.ScheduleDay, &tour.OpenTime, &tour.CloseTime, &tour.Price); err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 500)
			return
		}
		tours = append(tours, tour)
	}

	// Response
	var response models.ToursScheduleResponse

	if len(tours) == 0 {
		log.Println("(ERROR)\t", "Data empty")
		SendErrorResponse(w, 204)
		return
	} else {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = tours

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

		log.Println("(SUCCESS)\t", "Get tour list request")
	}
}
