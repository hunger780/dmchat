package services
const DBName ="DM_CHAT"
const Collection ="USERS"

func insertUser(){
	booksCollection := client.Database(DBName).Collection(Collection)
	var user DB_User
		insertResult, err := booksCollection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
}
