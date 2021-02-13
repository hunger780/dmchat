module main

go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	go.mongodb.org/mongo-driver v1.4.6
	services v1.0.0
)
replace services => ../dm/services
