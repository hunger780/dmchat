module github.com/hunger780/dmchat

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/hunger780/dmchat/dto v0.0.0-20210213172331-475c923a41b7 // indirect
	github.com/hunger780/dmchat/services v0.1.4
)

replace github.com/hunger780/dmchat/services => ./services
