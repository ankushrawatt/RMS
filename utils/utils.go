package utils

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Encode(writer http.ResponseWriter, i interface{}) {
	err := json.NewEncoder(writer).Encode(i)
	CheckError(err)
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
