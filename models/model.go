package models

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

type Hotel struct {
	ID            int    `json:"hotel_id"`
	HotelName     string `json:"hotel_name"`
	HotelStar     int    `json:"hotel_star"`
	HotelRating   int    `json:"hotel_rating"`
	HotelReview   int    `json:"hotel_review"`
	HotelFacility string `json:"hotel_facility"`
	HotelAddress  string `json:"hotel_address"`
	HotelCity     string `json:"hotel_city"`
	HotelCountry  string `json:"hotel_country"`
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

type Flight struct {
	ID                 int    `json:"flight_id"`
	AirplaneID         int    `json:"airplane_id"`
	DepartureAirport   int    `json:"departure_airport"`
	DestinationAirport int    `json:"destination_airport"`
	FlightType         string `json:"flight_type"`
	FlightNumber       string `json:"flight_number"`
	DepartureTime      string `json:"departure_time"`
	DepartureDate      string `json:"departure_date"`
	ArrivalTime        string `json:"arrival_time"`
	ArrivalDate        string `json:"arrival_date"`
	TravelTime         int    `json:"travel_time"`
}

type Bustrip struct {
	ID                    int    `json:"bustrip_id"`
	BusID                 int    `json:"bus_id"`
	DepartureBusstation   int    `json:"departure_busstation"`
	DestinationBusstation int    `json:"destination_busstation"`
	BusNumber             string `json:"bus_number"`
	DepartureTime         string `json:"departure_time"`
	DepartureDate         string `json:"departure_date"`
	ArrivalTime           string `json:"arrival_time"`
	ArrivalDate           string `json:"arrival_date"`
	TravelTime            int    `json:"travel_time"`
}

type Traintrip struct {
	ID                 int    `json:"traintrip_id"`
	TrainID            int    `json:"train_id"`
	DepartureStation   int    `json:"departure_station"`
	DestinationStation int    `json:"destination_station"`
	TraintripNumber    string `json:"trainTrip_number"`
	DepartureTime      string `json:"departure_time"`
	DepartureDate      string `json:"departure_date"`
	ArrivalTime        string `json:"arrival_time"`
	ArrivalDate        string `json:"arrival_date"`
	TravelTime         int    `json:"travel_time"`
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

type ToursSchedule struct {
	ID          int `json:"schedule_id"`
	TourName    int `json:"tour_id"`
	ScheduleDay int `json:"schedule_day"`
	OpenTime    int `json:"opentime"`
	CloseTime   int `json:"closetime"`
	Price       int `json:"price"`
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

type TrainResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    Traintrip `json:"data"`
}

type TrainsResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []Traintrip `json:"data"`
}
