package db

import (
	 "fmt"
	 "log"
 )
const dbName ="DM_CHAT"
const collection ="USERS"

func init(){
	usersCollection := Client.Database(dbName).Collection(collection)
}

func InsertUser(DB_User user){
	insertResult, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
    	log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return insertResult
}

func GetUserForUserID(userId string) (DB_User,string) {
    
	fmt.Println("Endpoint Hit: getUserForUserID")
    filter := Bson.D{{"User_ID", userId}}
	filterCursor, err := usersCollection.Find(ctx, filter)
	if err != nil {
    	log.Fatal(err)
	}
	var episodesFiltered []Bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
    	log.Fatal(err)
	}
	fmt.Println(episodesFiltered)
	
	
}
