package helper

import (
	"rmsProject/database"
	"rmsProject/model"
)

func Restaurant() ([]model.Restaurant, error) {
	SQL := `SELECT name,lat,lng,restaurantid from restaurant`
	data := make([]model.Restaurant, 0)
	err := database.RMS.Select(&data, SQL)
	if err != nil {
		return nil, err
	}
	//distance := make([]float64, 0)
	//fmt.Println(data[1].Lat)
	//for i := range data {
	//	distance[i] = utils.Distance(data[i].Lat, data[i].Lat, 28.7041, 77.1025)
	//}
	return data, nil
}

func AdminRestaurant(id string) ([]model.Restaurant, error) {
	SQL := `SELECT name,lat,lng,restaurantid from restaurant where createdby=$1`
	data := make([]model.Restaurant, 0)
	err := database.RMS.Select(&data, SQL, id)
	if err != nil {
		return nil, err
	}
	//for i := range data {
	//	data[i].Distance = utils.Distance(data[i].Lat, data[i].Lat, 28, 76)
	//	fmt.Println(data[1].Distance)
	//
	//}
	return data, nil
}

func AddRestaurant(name, createdBy string, lat, lng float64, restaurantID int) error {
	SQL := `INSERT INTO restaurant(name,lat,lng,restaurantid,createdby)VALUES($1,$2,$3,$4,$5) RETURNING name`
	var restaurantName string
	err := database.RMS.Get(&restaurantName, SQL, name, lat, lng, restaurantID, createdBy)
	if err != nil {
		return err
	}
	return nil
}

func AddDish(name string, ID, price int) error {
	SQL := `INSERT INTO dishes(restaurantid,dishname,price)VALUES($1,$2,$3) RETURNING dishname`
	var dish string
	err := database.RMS.Get(&dish, SQL, ID, name, price)
	if err != nil {
		return err
	}
	return nil
}

func Dishes(ID int) ([]model.Dishes, error) {
	SQL := `SELECT dishname,price FROM dishes where restaurantid=$1`
	dishes := make([]model.Dishes, 0)
	err := database.RMS.Select(&dishes, SQL, ID)
	if err != nil {
		return nil, err
	}
	return dishes, nil
}
func GetAdminID(userid string) (string, error) {
	//language=sql
	SQL := `SELECT id FROM users WHERE userid=$1`
	var id string
	err := database.RMS.Get(&id, SQL, userid)
	if err != nil {
		return "", err
	}
	return id, nil

}
