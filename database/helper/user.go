package helper

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"rmsProject/database"
	"rmsProject/model"
	"rmsProject/utils"
)

func HashPassword(password string) string {
	HashPass, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	utils.CheckError(err)
	return string(HashPass)
}

func CheckPassword(password, hashpass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(password))
	if err != nil {
		return errors.New("WRONG CREDENTIALS")
	}
	return nil

}

func CreateUser(ID, Email, Firstname, Lastname, UserID, Password, Mobile, role, createdby string) (string, error) {
	SQL := `INSERT INTO users(ID, Email, FirstName, LastName, UserID, Password, MobileNo,role,createdby) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)RETURNING UserID`
	var user string
	HashPass := HashPassword(Password)
	err := database.RMS.Get(&user, SQL, ID, Email, Firstname, Lastname, UserID, HashPass, Mobile, role, createdby)
	utils.CheckError(err)
	return user, nil
}

//
//func CreateAdmin(ID, email, firstname, lastname, password, mobile, userid, role string) (string, error) {
//	SQL := `INSERT INTO users(ID,Email,FirstName,LastName,UserID,Password,MobileNo,role)VALUES($1,$2,$3,$4,$5,$6,$7,$8)returning UserID`
//	var admin string
//	HashPass := HashPassword(password)
//	err := database.RMS.Get(&admin, SQL, ID, email, firstname, lastname, userid, HashPass, mobile, role)
//	utils.CheckError(err)
//	return admin, nil
//}

//func CreateSubAdmin(ID, email, firstname, lastname, password, mobile, userid, role string) (string, error) {
//	SQL := `INSERT INTO users(ID,Email,FirstName,LastName,UserID,Password,MobileNo,role)VALUES($1,$2,$3,$4,$5,$6,$7,$8)returning UserID`
//	var admin string
//	HashPass := HashPassword(password)
//	err := database.RMS.Get(&admin, SQL, ID, email, firstname, lastname, userid, HashPass, mobile, role)
//	utils.CheckError(err)
//	return admin, nil
//}

func LoginUser(email, password string) (*model.Info, error) {
	SQL := `SELECT userid,password,role FROM users WHERE email=$1`
	var hashPass string
	var info model.Info
	err := database.RMS.QueryRowx(SQL, email).Scan(&info.Userid, &hashPass, &info.Role)
	if err != nil {
		return nil, err
	}
	passErr := CheckPassword(password, hashPass)
	if passErr != nil {
		return nil, passErr
	}
	return &info, nil
}

func CreateSession(token, userid string) error {
	SQL := `INSERT INTO session(id,userid)VALUES($1,$2) returning userid`
	var user string
	err := database.RMS.Get(&user, SQL, token, userid)
	if err != nil {
		return err
	}
	return nil
}
