module main

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	go.mongodb.org/mongo-driver v1.4.6
	github.com/hunger780/dmchat/dm/services v0.1.0
)
replace github.com/hunger780/dmchat/dm/services => ../dm/services
