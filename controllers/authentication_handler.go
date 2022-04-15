package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Get value from env
var stringKey = LoadEnv("JWT_KEY")
var jwtKey = []byte(stringKey)
var tokenName = "loginToken"

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	UserType int    `json:"user_type"`
	jwt.StandardClaims
}

func GenerateToken(w http.ResponseWriter, id int, username string, userType int) {
	tokenExpiryTime := time.Now().Add(10 * time.Minute)

	// create claims with user data
	claims := &Claims{
		ID:       id,
		Username: username,
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

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := ValidateUserToken(r, accessType)
		if !isValidToken {
			SendErrorResponse(w, 401)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func ValidateUserToken(r *http.Request, accessType int) bool {
	isAccessTokenValid, id, username, userType := ValidateTokenFormCookies(r)
	fmt.Println(id, username, userType, accessType, isAccessTokenValid)

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
			return true, accessClaims.ID, accessClaims.Username, accessClaims.UserType
		} else {
			log.Println(err2)
		}
	} else {
		log.Println(err1)
	}
	return false, -1, "", -1
}

func GetIdFromCookie(r *http.Request) int {

	cookie, err1 := r.Cookie(tokenName)
	if err1 == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err2 := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err2 == nil && parsedToken.Valid {
			return accessClaims.ID
		} else {
			log.Println(err2)
		}
	} else {
		log.Println(err1)
	}
	return -1
}
