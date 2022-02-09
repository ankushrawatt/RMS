package handler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"math"
	"net/http"
	"rmsProject/database/helper"
	"rmsProject/model"
	"rmsProject/utils"
	"time"
)

var mySigningKey = []byte("secret_key")

func CreateToken(userid, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = userid
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func Distance(lat1, lng1, lat2, lng2 float64) float64 {
	const PI float64 = 3.1415926535
	radLat1 := float64(PI * lat1 / 180)
	radLat2 := float64(PI * lat2 / 180)
	theta := float64(lng1 - lng2)
	dist := math.Sin(radLat1)*math.Sin(radLat2) + math.Cos(radLat1)*math.Cos(radLat2)*math.Cos(theta)
	if dist > 1 {
		dist = 1
	}
	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515
	dist = dist * 1.609344 //distance in km

	return dist
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	var info model.UserInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	utils.CheckError(err)
	ID := uuid.New()
	user, NewErr := helper.CreateUser(ID.String(), info.Email, info.FirstName, info.LastName, info.UserID, info.Password, info.MobileNo, info.Role)
	utils.CheckError(NewErr)
	utils.Encode(writer, user)
}

//func AdminSignUp(writer http.ResponseWriter, request *http.Request) {
//	var info model.UserInfo
//	role := "admin"
//	err := json.NewDecoder(request.Body).Decode(&info)
//	utils.CheckError(err)
//	ID := uuid.New()
//	admin, NewErr := helper.CreateUser(ID.String(), info.Email, info.FirstName, info.LastName, info.UserID, info.Password, info.MobileNo, role)
//	utils.CheckError(NewErr)
//	utils.Encode(writer, admin)
//
//}

func Login(writer http.ResponseWriter, request *http.Request) {
	var info model.LoginInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	utils.CheckError(err)
	userid, loginErr := helper.LoginUser(info.Email, info.Password)
	utils.CheckError(loginErr)
	token, tokenErr := CreateToken(userid.Userid, userid.Role)
	utils.CheckError(tokenErr)
	sessionErr := helper.CreateSession(token, userid.Userid)
	utils.CheckError(sessionErr)
	err = json.NewEncoder(writer).Encode(token)
	utils.CheckError(err)
}

func AddRestaurant(writer http.ResponseWriter, request *http.Request) {
	var restaurant model.Restaurant
	err := json.NewDecoder(request.Body).Decode(&restaurant)
	utils.CheckError(err)
	helper.AddRestaurant(restaurant.Name, restaurant.Lat, restaurant.Lng)

}

func AllRestaurant(writer http.ResponseWriter, request *http.Request) {
	data, err := helper.Restaurant()
	utils.CheckError(err)
	err = json.NewEncoder(writer).Encode(data)
	utils.CheckError(err)
}

func AddSubAdmin(writer http.ResponseWriter, request *http.Request) {
	var info model.UserInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	utils.CheckError(err)
	ID := uuid.New()
	userID, NewErr := helper.CreateUser(ID.String(), info.Email, info.FirstName, info.LastName, info.UserID, info.Password, info.MobileNo, info.Role)
	utils.CheckError(NewErr)
	err = json.NewEncoder(writer).Encode(userID)
	utils.CheckError(err)
}
