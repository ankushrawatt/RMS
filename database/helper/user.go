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

func Subadmin(admin, role string) ([]model.SubAdmin, error) {
	// language=SQL
	SQL := `SELECT email,userid,mobileno FROM users WHERE createdby=$1 AND role=$2 `
	subadmin := make([]model.SubAdmin, 0)
	err := database.RMS.Select(&subadmin, SQL, admin, role)
	if err != nil {
		return nil, err
	}
	return subadmin, nil
}

func UsersByAdmin(adminID string) ([]model.UserInfo, error) {
	// language=SQL
	SQL := `SELECT id,firstname,lastname,email,mobileno,role,userid FROM users WHERE createdby=$1`
	users := make([]model.UserInfo, 0)
	err := database.RMS.Select(&users, SQL, adminID)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func AddAddress(id string, lat, lng float64) error {
	//language=SQL
	SQL := `INSERT INTO address(id, lat, lng) VALUES ($1,$2,$3) returning id`
	var userid string
	err := database.RMS.Get(&userid, SQL, id, lat, lng)
	if err != nil {
		return err
	}
	return nil
}

func GetUserAddress(id string) (*model.Address, error) {
	//language=sql
	SQL := `SELECT lat,lng from address where id=$1`
	var address model.Address
	err := database.RMS.Get(&address, SQL, id)
	if err != nil {
		return nil, err
	}
	return &address, nil

}
