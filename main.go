package main

import (
   // "encoding/json"
    "fmt"
    "log"
  //  "io/ioutil"
    "net/http"
	"github.com/hunger780/dmchat/dm/services"
   
    "github.com/gorilla/mux"
)



func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}



func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/signup", services.Signup).Methods("POST")
    myRouter.HandleFunc("/users", services.ReturnAllUsers)
    myRouter.HandleFunc("/signin", services.Signin).Methods("POST")
    services.Connect() 
    log.Fatal(http.ListenAndServe(":8080", myRouter))
    
}

func main() {

    handleRequests()
}
