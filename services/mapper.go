package services

import(
	//"fmt"
	"github.com/hunger780/dmchat/dto"
	"github.com/hunger780/dmchat/db"
)
func MapUserToDBUser(user dto.User) db.DB_User {
	
	return &db.DB_User{
		User_ID: user.userid,
		First_Name: user.firstName,
		Last_Name: user.lastName,
		DOB : user.dob,
		Password : user.password}

}
