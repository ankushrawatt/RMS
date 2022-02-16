package model

import "github.com/dgrijalva/jwt-go"

type UserInfo struct {
	FirstName string `db:"firstname" json:"firstName"`
	LastName  string `db:"lastname" json:"lastName"`
	UserID    string `db:"userid" json:"userID"`
	MobileNo  string `db:"mobileno" json:"mobileNo"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	ID        string `db:"id" json:"ID"`
	Role      string `db:"role" json:"role"`
}

type JWTClaims struct {
	UserID string `db:"userid"json:"user"`
	Role   string `db:"role" json:"role"`
	jwt.StandardClaims
}

type LoginInfo struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type Info struct {
	Userid string `db:"userid"`
	Role   string `db:"role"`
}

type Address struct {
	Lat float64 `db:"lat" json:"lat"`
	Lng float64 `db:"lng" json:"lng"`
}
