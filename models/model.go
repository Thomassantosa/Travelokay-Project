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
