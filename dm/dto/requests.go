package dto
type DB_User struct {
    Id      string    `json:"id"`
	User_ID   string `json:"userid"`
    First_Name   string `json:"firstName"`
    Last_Name    string `json:"lastName"`
    DOB string `json:"dob"`
	Password string `json:"password"`
}