package model

type AdminInfo struct {
	FirstName string `db:"firstname" json:"firstName"`
	LastName  string `db:"lastname" json:"lastName"`
	MobileNo  string `db:"mobileno" json:"mobileNo"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	ID        string `db:"id" json:"ID"`
	AdminID   string `db:"adminid" json:"adminID"`
}

type SubAdmin struct {
	//FirstName string `db:"firstname" json:"firstName"`
	//LastName  string `db:"lastname" json:"lastName"`
	UserID   string `db:"userID" json:"userID"`
	MobileNo string `db:"mobileno" json:"mobileNo"`
	Email    string `db:"email" json:"email"`
}
