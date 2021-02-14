package db

import "github.com/zebresel-com/mongodm"
import "io/ioutil"
import "fmt"
import "os"
import "encoding/json"

func Connect(){
	file, err := ioutil.ReadFile("locals.json")

	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var localMap map[string]map[string]string
	json.Unmarshal(file, &localMap)

	dbConfig := &mongodm.Config{
		DatabaseHosts: []string{"cluster0.sy3jf.mongodb.net"},
		DatabaseName: "DM_CHAT",
		DatabaseUser: "rakesh",
		DatabasePassword: "10980000gG#",
		// The option `DatabaseSource` is the database used to establish 
		// credentials and privileges with a MongoDB server. Defaults to the value
		// of `DatabaseName`, if that is set, or "admin" otherwise.
		DatabaseSource: "DM_CHAT",
		Locals:       localMap["en-US"],
	}

	connection, err := mongodm.Connect(dbConfig)

	if err != nil {
		fmt.Println("Database connection error: %v", err)
	}
	// See https://godoc.org/github.com/zebresel-com/mongodm#Connection.Register
	connection.Register(&model.User{}, "users")
}

