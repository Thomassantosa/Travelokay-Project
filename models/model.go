package models

type User struct {
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
	ID            int     `json:"id"`
	HotelName     string  `json:"hotel_name"`
	HotelStar     int     `json:"hotel_star"`
	HotelRating   float32 `json:"hotel_rating"`
	HotelReview   int     `json:"hotel_review"`
	HotelFacility string  `json:"hotel_facility"`
	HotelAddress  string  `json:"hotel_address"`
	HotelCity     string  `json:"hotel_city"`
	HotelCountry  string  `json:"hotel_country"`
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
