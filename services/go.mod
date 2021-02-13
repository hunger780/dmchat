module github.com/hunger780/dmchat/services

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/hunger780/dmchat/db v0.0.0-20210213180306-230647087863
	github.com/hunger780/dmchat/dto v0.0.0-20210213172331-475c923a41b7
	go.mongodb.org/mongo-driver v1.4.6
//github.com/hunger780/dmchat/dto v0.1.4
)

//replace github.com/hunger780/dmchat => ./dmchat
