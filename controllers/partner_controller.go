package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Travelokay-Project/models"
)

func GetFlightPartnerList(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	userId := GetIdFromCookie(r)
	userCompanyName := ""

	//Get Data Company Name Partner
	queryGetCompanyName := "SELECT company_name FROM users WHERE user_id = ?"
	rows, errQuery := db.Query(queryGetCompanyName, userId)
	if errQuery != nil {
		SendErrorResponse(w, 500)
		log.Println(errQuery)
		return
	}
	for rows.Next() {
		err := rows.Scan(&userCompanyName)
		if err != nil {
			SendErrorResponse(w, 500)
			log.Println(err)
			return
		}
	}

	queryGetListFlights :=
		`SELECT flights.flight_id, airplanes.airplane_model, airlines.airline_name, airportA.airport_id, airportA.airport_code,` +
			` airportA.airport_name, airportA.airport_city, airportA.airport_country, airportB.airport_id, airportB.airport_code,` +
			` airportB.airport_name, airportB.airport_city, airportB.airport_country, flight_type, flight_number, departure_time,` +
			` arrival_time, travel_time FROM flights` +
			` JOIN airplanes ON flights.airplane_id = airplanes.airplane_id` +
			` JOIN airlines ON airplanes.airline_id = airlines.airline_id` +
			` JOIN airports AS airportA ON flights.departure_airport = airportA.airport_id` +
			` JOIN airports AS airportB ON flights.destination_airport = airportB.airport_id` +
			` WHERE airlines.airline_name = ?` +
			` GROUP BY flights.flight_id`

	rowsFlights, errQueryFlights := db.Query(queryGetListFlights, userCompanyName)

	log.Println(queryGetListFlights)
	if errQueryFlights != nil {
		SendErrorResponse(w, 500)
		log.Println(errQueryFlights)
		return
	}

	var flight models.Flight
	var flights []models.Flight

	for rowsFlights.Next() {
		err := rowsFlights.Scan(&flight.ID, &flight.AirplaneModel, &flight.AirlineName, &flight.DepartureAirport.ID, &flight.DepartureAirport.Code,
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

	log.Println("(SUCCESS)\t", "Get list flights request")
}

func AddNewFlight(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(w, 500)
		log.Println(err)
		return
	}

	airplaneId := r.Form.Get("airplaneId")
	departureAirport := r.Form.Get("departureAirport")
	destinationAirport := r.Form.Get("destinationAirport")
	flightType := r.Form.Get("flightType")
	flightNumber := r.Form.Get("flightNumber")
	departureTime := r.Form.Get("departureTime")
	arrivalTime := r.Form.Get("arrivalTime")
	travelTime := r.Form.Get("travelTime")

	query := `
		INSERT INTO flights(airplane_id, departure_airport, 
		destination_airport, flight_type, flight_number, departure_time, 
		arrival_time, travel_time) VALUES (?,?,?,?,?,?,?,?)
	`

	_, errQuery := db.Exec(query, airplaneId, departureAirport, destinationAirport, flightType, flightNumber, departureTime, arrivalTime, travelTime)
	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
		return
	} else {
		log.Println("(SUCCESS)\t", "Add new flight request")
		SendSuccessResponse(w)
		return
	}

}

func UpdatePartner(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from cookie
	partnerId := GetIdFromCookie(r)

	// Get value from form
	fullname := r.FormValue("fullname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	address := r.FormValue("address")
	partnerType := r.FormValue("partnerType")
	companyName := r.FormValue("companyName")

	// encrypt password
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	// Query
	query := "UPDATE users SET"

	if fullname != "" {
		query += " fullname = '" + fullname + "',"
	}
	if username != "" {
		query += " username = '" + username + "',"
	}
	if email != "" {
		query += " email = '" + email + "',"
	}
	if encryptedPassword != "" {
		query += " password = '" + encryptedPassword + "',"
	}
	if address != "" {
		query += " address = '" + address + "',"
	}
	if partnerType != "" {
		query += " address = '" + partnerType + "',"
	}
	if companyName != "" {
		query += " address = '" + companyName + "',"
	}

	queryNew := query[:len(query)-1] // Delete last coma
	queryNew += " WHERE user_id = " + strconv.Itoa(partnerId)

	_, errQuery := db.Exec(queryNew)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
	} else {
		SendSuccessResponse(w)
		log.Println("(SUCCESS)\t", "Update partner request")
	}
}

func DeleteFlight(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(w, 500)
		log.Println(err)
		return
	}

	//Pakai Params
	flightId := r.FormValue("flightId")

	_, errQuery := db.Exec("DELETE FROM flights WHERE flight_id = ?", flightId)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
	} else {
		log.Println("(SUCCESS)\t", "Delete flight request")
		SendSuccessResponse(w)
	}
}

func AddNewAirline(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(w, 500)
		log.Println("(ERROR)\t", err)
		return
	}

	airlineName := r.Form.Get("airlineName")
	airlineContact := r.Form.Get("airlineContact")

	query := "INSERT INTO airlines(airline_name, airline_contact) VALUES (?,?)"

	_, errQuery := db.Exec(query, airlineName, airlineContact)
	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
		return
	} else {
		log.Println("(SUCCESS)\t", "Add new airline request")
		SendSuccessResponse(w)
		return
	}

}

func AddNewAirplane(w http.ResponseWriter, r *http.Request) {

	// Connect to database
	db := Connect()
	defer db.Close()

	// Get value from form
	err := r.ParseForm()
	if err != nil {
		SendErrorResponse(w, 500)
		log.Println("(ERROR)\t", err)
		return
	}

	airlineId := r.Form.Get("airlineId")
	airplaneModel := r.Form.Get("airplaneModel")

	//Check if airline is already
	if !CheckAirlineAlready(w, r, airlineId) {
		log.Println("Data airplane tidak ada")
		return
	} else {

		log.Println("Data airplane ada")
	}

	query := "INSERT INTO airplanes(airline_id, airplane_model) VALUES (?,?)"

	_, errQuery := db.Exec(query, airlineId, airplaneModel)
	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 400)
		return
	} else {
		log.Println("(SUCCESS)\t", "Add new airline request")
		SendSuccessResponse(w)
		return
	}

}

func CheckAirlineAlready(w http.ResponseWriter, r *http.Request, airlineId string) bool {

	db := Connect()
	defer db.Close()

	// Query
	rows, errQuery := db.Query("SELECT airline_id FROM airplanes WHERE airline_id = ?", airlineId)

	if errQuery != nil {
		log.Println("(ERROR)\t", errQuery)
		SendErrorResponse(w, 500)
		return false
	}

	var airline int

	for rows.Next() {
		if err := rows.Scan(&airline); err != nil {
			log.Println("(ERROR)\t", err)
			SendErrorResponse(w, 500)
			return false
		}
	}

	return true
}
