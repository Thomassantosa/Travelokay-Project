package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	models "github.com/Travelokay-Project/models"
	"github.com/golang-jwt/jwt/v4"
)

var stringKey = LoadEnv("JWT_KEY")
var jwtKey = []byte(stringKey)
var tokenName = "token"

type Claims struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserType int    `json:"user_type"`
	jwt.StandardClaims
}

func GenerateToken(w http.ResponseWriter, id int, name string, userType int) {
	tokenExpiryTime := time.Now().Add(5 * time.Minute)

	// create claims with user data
	claims := &Claims{
		ID:       id,
		Name:     name,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	// encrypt claims to jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	// set token to cookie
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})
}

func ResetUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func SendUnAuthorizedResponse(w http.ResponseWriter) {
	// success response
	var response models.MessageResponse
	response.Status = 401
	response.Message = "UnAuthorized Access"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := ValidateUserToken(r, accessType)
		if !isValidToken {
			SendUnAuthorizedResponse(w)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func ValidateUserToken(r *http.Request, accessType int) bool {
	isAccessTokenValid, id, email, userType := ValidateTokenFormCookies(r)
	fmt.Println(id, email, userType, accessType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}
	return false
}

func ValidateTokenFormCookies(r *http.Request) (bool, int, string, int) {

	cookie, err1 := r.Cookie(tokenName)
	log.Println(cookie)
	if err1 == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err2 := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err2 == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.UserType
		} else {
			log.Println(err2)
		}
	} else {
		log.Println(err1)
	}
	return false, -1, "", -1
}
