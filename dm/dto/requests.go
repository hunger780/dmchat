package github.com/hunger780/dmchat/dm/dto

type User struct {
    id      string    `json:"id"`
	userid   string `json:"userid"`
    firstName   string `json:"firstName"`
    lastName    string `json:"lastName"`
    dob string `json:"dob"`
	password string `json:"password"`
}