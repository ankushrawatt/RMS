package model

type Distance struct {
	Name         string  `db:"name" json:"name"`
	Lat          float64 `db:"lat" json:"lat"`
	Lng          float64 `db:"lng" json:"lng"`
	RestaurantID int     `db:"restaurantid" json:"restaurantID"`
	Distance     float64 `json:"distance"`
}

type Restaurant struct {
	Name         string  `db:"name" json:"name"`
	Lat          float64 `db:"lat" json:"lat"`
	Lng          float64 `db:"lng" json:"lng"`
	RestaurantID int     `db:"restaurantid" json:"restaurantID"`
}

type Dishes struct {
	ID       int    `db:"restaurantid" json:"ID"`
	DishName string `db:"dishname" json:"dishName"`
	Price    int    `db:"price" json:"price"`
}
