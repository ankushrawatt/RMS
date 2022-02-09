package model

import "github.com/dgrijalva/jwt-go"

type UserInfo struct {
	FirstName string `db:"FirstName" json:"firstName"`
	LastName  string `db:"LastName" json:"lastName"`
	UserID    string `db:"UserID" json:"userID"`
	MobileNo  string `db:"MobileNo" json:"mobileNo"`
	Email     string `db:"Email" json:"email"`
	Password  string `db:"Password" json:"password"`
	ID        string `db:"ID" json:"ID"`
	Role      string `db:"role" json:"role"`
}

type JWTClaims struct {
	UserID int    `json:"user"`
	role   string `json:"role"`
	jwt.StandardClaims
}

type AdminInfo struct {
	FirstName string `db:"FirstName" json:"firstName"`
	LastName  string `db:"LastName" json:"lastName"`
	MobileNo  string `db:"MobileNo" json:"mobileNo"`
	Email     string `db:"Email" json:"email"`
	Password  string `db:"Password" json:"password"`
	ID        string `db:"ID" json:"ID"`
	AdminID   string `db:"AdminID" json:"adminID"`
}

type Restaurant struct {
	Name         string  `db:"name" json:"name"`
	Lat          float64 `db:"LAT" json:"lat"`
	Lng          float64 `db:"LNG" json:"lng"`
	RestaurantID int     `db:"restaurantID" json:"restaurantID"`
}

type LoginInfo struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type Info struct {
	Userid string `db:"userid"`
	Role   string `db:"role"`
}
