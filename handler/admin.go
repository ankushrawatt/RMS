package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"rmsProject/database/helper"
	"rmsProject/model"
	"rmsProject/utils"
)

func GetID(userid string) string {
	id, err := helper.GetAdminID(userid)
	utils.CheckError(err)
	return id
}

func AddRestaurant(writer http.ResponseWriter, request *http.Request) {
	var restaurant model.Restaurant
	err := json.NewDecoder(request.Body).Decode(&restaurant)
	utils.CheckError(err)
	//fmt.Println(request.Context().Value("user"))

	claims := request.Context().Value("user").(model.JWTClaims)
	id := GetID(claims.UserID)
	NewErr := helper.AddRestaurant(restaurant.Name, id, restaurant.Lat, restaurant.Lng, restaurant.RestaurantID)
	utils.CheckError(NewErr)
}

//For user
func AllRestaurant(writer http.ResponseWriter, request *http.Request) {
	//claims := request.Context().Value("user").(model.JWTClaims)
	//id:=GetID(claims.UserID)
	data, err := helper.Restaurant()

	//fmt.Println(claims.Role)
	utils.CheckError(err)
	err = json.NewEncoder(writer).Encode(data)
	utils.CheckError(err)
}

func AddDish(writer http.ResponseWriter, request *http.Request) {
	var dish model.Dishes
	err := json.NewDecoder(request.Body).Decode(&dish)
	utils.CheckError(err)
	//claims := request.Context().Value("user").(model.JWTClaims)
	//id := GetID(claims.UserID)
	err = helper.AddDish(dish.DishName, dish.ID, dish.Price)
	utils.CheckError(err)
}

//For user
func AllDish(writer http.ResponseWriter, request *http.Request) {
	//claims := request.Context().Value("user").(model.JWTClaims)
	//id := GetID(claims.UserID)
	id := struct {
		id int `json:"ID"`
	}{}
	json.NewDecoder(request.Body).Decode(&id)
	dishes, err := helper.Dishes(id.id)
	utils.CheckError(err)
	err = json.NewEncoder(writer).Encode(dishes)
	utils.CheckError(err)
}

func AddSubAdmin(writer http.ResponseWriter, request *http.Request) {
	var info model.UserInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	utils.CheckError(err)
	role := "sub-admin"
	claims := request.Context().Value("user").(model.JWTClaims)
	adminID := GetID(claims.UserID)
	ID := uuid.New()
	userID, NewErr := helper.CreateUser(ID.String(), info.Email, info.FirstName, info.LastName, info.UserID, info.Password, info.MobileNo, role, adminID)
	utils.CheckError(NewErr)
	err = json.NewEncoder(writer).Encode(userID)
	utils.CheckError(err)
}

func Subadmin(writer http.ResponseWriter, request *http.Request) {
	claims := request.Context().Value("user").(model.JWTClaims)
	role := "sub-admin"
	adminID := GetID(claims.UserID)
	subadmin, err := helper.Subadmin(adminID, role)
	utils.CheckError(err)
	err = json.NewEncoder(writer).Encode(subadmin)
	utils.CheckError(err)
}

func AdminUsers(writer http.ResponseWriter, request *http.Request) {
	claims := request.Context().Value("user").(model.JWTClaims)
	ID := GetID(claims.UserID)
	users, err := helper.AdminUsers(ID)
	utils.CheckError(err)
	err = json.NewEncoder(writer).Encode(users)
	utils.CheckError(err)
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	var info model.UserInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	utils.CheckError(err)
	ID := uuid.New()
	role := "user"
	claims := request.Context().Value("user").(model.JWTClaims)
	userID, NewErr := helper.CreateUser(ID.String(), info.Email, info.FirstName, info.LastName, info.UserID, info.Password, info.MobileNo, role, GetID(claims.UserID))
	utils.CheckError(NewErr)
	err = json.NewEncoder(writer).Encode(userID)
	utils.CheckError(err)
}

func AdminRestaurant(writer http.ResponseWriter, request *http.Request) {
	claims := request.Context().Value("user").(model.JWTClaims)
	adminID := GetID(claims.UserID)
	data, err := helper.AdminRestaurant(adminID)
	utils.CheckError(err)
	err = json.NewEncoder(writer).Encode(data)
	utils.CheckError(err)
}
