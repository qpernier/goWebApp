package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       int    `bson:"id,omitempty"`
	Username string `bson:"username,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

func (user User) String() string {
	return fmt.Sprintf("{Id : %d, Username : %s, Email : %s, Password: %s}", user.Id, user.Username, user.Email, user.Password)
}

var usersMock = make(map[string]User)
var client *mongo.Client
var ctx context.Context
var collection *mongo.Collection

func init() {
	usersMock["bob"] = User{Id: 1,
		Username: "bob",
		Password: "bob",
		Email:    "bob@gmail.com"}
	usersMock["alice"] = User{Id: 2,
		Username: "alice",
		Password: "alice",
		Email:    "alice@gmail.com"}

	/*
		Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	/*
	   Get my collection instance
	*/
	collection := client.Database("test").Collection("users")

	defer client.Disconnect(ctx)

	/*
		List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Databases XXXXXXXXXXXXXXXXXXXXXXxxxxxxxx")
	fmt.Println(databases)

	/*
	   Iterate a cursor and print it
	*/
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var users []User
	if err = cur.All(ctx, &users); err != nil {
		panic(err)
	}
	fmt.Println("USERS XXXXXXXXXXXXXXXXXXXXXxxxxxxxxxx")
	fmt.Println(users)

}

func GetUser(username string) User {
	return usersMock[username]
}
