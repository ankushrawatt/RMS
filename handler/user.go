package handler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
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
	claims["user"] = userid //shiv
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	utils.CheckError(err)
	return tokenString, nil
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	var info model.UserInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	utils.CheckError(err)
	ID := uuid.New()
	user, NewErr := helper.CreateUser(ID.String(), info.Email, info.FirstName, info.LastName, info.UserID, info.Password, info.MobileNo, info.Role, "SELF")
	utils.CheckError(NewErr)
	utils.Encode(writer, user)
}

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

func Health(writer http.ResponseWriter, request *http.Request) {

	err := json.NewEncoder(writer).Encode("Hello")
	utils.CheckError(err)
}

func AddAddress(writer http.ResponseWriter, request *http.Request) {
	claims := request.Context().Value("user").(model.JWTClaims)
	ID := GetID(claims.UserID)
	var address model.Address
	err := json.NewDecoder(request.Body).Decode(&address)
	utils.CheckError(err)
	err = helper.AddAddress(ID, address.Lat, address.Lng)
	utils.CheckError(err)

}
