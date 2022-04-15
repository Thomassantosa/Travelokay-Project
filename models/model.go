package models

import "database/sql"

type User struct {
	ID          int    `json:"id"`
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	UserType    int    `json:"user_type"`
	DateCreated string `json:"date_created"`
}

type Partner struct {
	ID          int    `json:"id"`
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	UserType    int    `json:"user_type"`
	PartnerType string `json:"partner_type"`
	CompanyName string `json:"company_name"`
	DateCreated string `json:"date_created"`
}

type Flight struct {
	ID                 int     `json:"flight_id"`
	AirplaneModel      string  `json:"airplane_model"`
	AirlineName        string  `json:"airline_name"`
	DepartureAirport   Airport `json:"departure_airport"`
	DestinationAirport Airport `json:"destination_airport"`
	FlightType         string  `json:"flight_type"`
	FlightNumber       string  `json:"flight_number"`
	DepartureTime      string  `json:"departure_time"`
	ArrivalTime        string  `json:"arrival_time"`
	TravelTime         int     `json:"travel_time"`
}

type Airport struct {
	ID      int    `json:"airport_id"`
	Code    string `json:"airport_code"`
	Name    string `json:"airport_name"`
	City    string `json:"airport_city"`
	Country string `json:"airport_country"`
}

type Traintrip struct {
	ID                 int     `json:"traintrip_id"`
	TrainModel         string  `json:"train_model"`
	DepartureStation   Station `json:"departure_station"`
	DestinationStation Station `json:"destination_station"`
	TrainTripNumber    string  `json:"traintrip_number"`
	DepartureTime      string  `json:"departure_time"`
	ArrivalTime        string  `json:"arrival_time"`
	TravelTime         int     `json:"travel_time"`
}

type Station struct {
	ID   int    `json:"station_id"`
	Code string `json:"station_code"`
	Name string `json:"station_name"`
	City string `json:"station_city"`
}

type Bustrip struct {
	ID                    int        `json:"bustrip_id"`
	BusModel              string     `json:"bus_model"`
	CompanyName           string     `json:"buscompany_name"`
	DepartureBusstation   Busstation `json:"departure_busstation"`
	DestinationBusstation Busstation `json:"destination_busstation"`
	BustripNumber         string     `json:"bustrip_number"`
	DepartureTime         string     `json:"departure_time"`
	ArrivalTime           string     `json:"arrival_time"`
	TravelTime            int        `json:"travel_time"`
}

type Busstation struct {
	ID   int    `json:"busstation_id"`
	Code string `json:"busstation_code"`
	Name string `json:"busstation_name"`
	City string `json:"busstation_city"`
}

type Seat struct {
	ID              int    `json:"seat_id"`
	SeatType        string `json:"seat_type"`
	SeatName        string `json:"seat_name"`
	SeatStatus      int    `json:"seat_status"`
	BaggageCapacity int    `json:"baggage_capacity"`
	SeatPrice       int    `json:"seat_price"`
}

type Hotel struct {
	ID            int     `json:"hotel_id"`
	HotelName     string  `json:"hotel_name"`
	HotelStar     int     `json:"hotel_star"`
	HotelRating   float32 `json:"hotel_rating"`
	HotelReview   int     `json:"hotel_review"`
	HotelFacility string  `json:"hotel_facility"`
	HotelAddress  string  `json:"hotel_address"`
	HotelCity     string  `json:"hotel_city"`
	HotelCountry  string  `json:"hotel_country"`
}

type Room struct {
	ID           int    `json:"room_id"`
	HotelID      int    `json:"hotel_id"`
	RoomName     string `json:"room_name"`
	RoomType     string `json:"room_type"`
	RoomPrice    int    `json:"room_price"`
	RoomFacility string `json:"room_facility"`
	RoomCapacity int    `json:"room_capacity"`
	RoomStatus   int    `json:"room_status"`
	CheckIn      string `json:"checkin"`
	CheckOut     string `json:"checkout"`
}

type Tours struct {
	ID           int     `json:"tour_id"`
	TourName     string  `json:"tour_name"`
	TourRating   float32 `json:"tour_rating"`
	TourReview   int     `json:"tour_review"`
	TourDesc     string  `json:"tour_desc"`
	TourFacility string  `json:"tour_facility"`
	TourAddress  string  `json:"tour_address"`
	TourCity     string  `json:"tour_city"`
	TourProvince string  `json:"tour_province"`
	TourCountry  string  `json:"tour_country"`
}

type Orders struct {
	ID              int           `json:"order_id"`
	UserID          int           `json:"user_id"`
	SeatID          sql.NullInt64 `json:"seat_id"`
	RoomID          sql.NullInt64 `json:"room_id"`
	TourScheduleID  sql.NullInt64 `json:"tourschedule_id"`
	OrderDate       string        `json:"order_date"`
	OrderStatus     string        `json:"order_status"`
	TransactionType string        `json:"transaction_type"`
}

type ToursSchedule struct {
	ID          int    `json:"schedule_id"`
	TourID      int    `json:"tour_id"`
	ScheduleDay int    `json:"schedule_day"`
	OpenTime    string `json:"opentime"`
	CloseTime   string `json:"closetime"`
	Price       int    `json:"price"`
}

type MessageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type PartnerResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Partner `json:"data"`
}

type PartnersResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Partner `json:"data"`
}

type HotelResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Hotel  `json:"data"`
}

type HotelsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Hotel `json:"data"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type FlightResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Flight `json:"data"`
}

type FlightsResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Flight `json:"data"`
}
type BusResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Bustrip `json:"data"`
}

type BusesResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Bustrip `json:"data"`
}

type TrainTripsResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []Traintrip `json:"data"`
}

type ToursResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Tours `json:"data"`
}

type ToursScheduleResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []ToursSchedule `json:"data"`
}

type SeatsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Seat `json:"data"`
}
type OrdersResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Orders `json:"data"`
}
