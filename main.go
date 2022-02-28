package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Ref      string             `json:"ref" bson:"ref"`
	Name     string             `json:"user_name" bson:"user_name"`
	Team     string             `json:"team_name" bson:"team_name"`
	Location string             `json:"location_name" bson:"location_name"`
}

func mongoconnection() (context.Context, *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("testDB").Collection("users")
	defer cancel()
	return ctx, collection
}

// function for handling get request for single element
func getreq(c echo.Context) error {
	//function to get data and handle get requests
	var refno string = c.QueryParam("ref")
	var result User
	// initiate and establish connection
	_, collection := mongoconnection()
	// find one
	filter := bson.D{{"ref", refno}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, map[string]User{
		"data": result,
	})
}

// function for handling get request for single element
func getreqmul(c echo.Context) error {
	// initiate and establish connection
	_, collection := mongoconnection()
	// find one
	results, err := collection.Distinct(context.TODO(), "user_name", bson.D{})
	if err != nil {
	}
	return c.JSON(http.StatusOK, map[string][]interface{}{
		"data": results,
	})
}

// function for put request
func putreq(c echo.Context) error {
	//function to get data and handle get requests
	var refno string = c.QueryParam("ref")
	username := c.FormValue("user_name")
	fmt.Printf("username")
	fmt.Printf(username)
	// var username string = c.FormValue("user_name")
	// var userlocation string = c.FormValue("location_name")
	// var userteam string = c.FormValue("team_name")
	// initiate and eestablish connection
	_, collection := mongoconnection()
	// find and put the data
	filter := bson.D{{"ref", refno}}
	update := bson.D{{"$set", bson.D{{"user_name", "username"}, {"team_name", "userteam"}, {"location_name", "userlocation"}}}}
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "updated the values for the reference id",
	})
}

// delete the request
func deletereq(c echo.Context) error {
	//function to get data and handle get requests
	var refno string = c.FormValue("ref")
	// initiate and eestablish connection
	_, collection := mongoconnection()
	// find and delete the reference id
	filter := bson.D{{"ref", refno}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted the value for the given user"})
}

// main function
func main() {
	e := echo.New()
	//middleware cors defined
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5500", "http://localhost:8080", "http://localhost:8081"},
		AllowMethods: []string{"GET", "PUT", "DELETE"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	//Get Data
	e.GET("/", getreqmul)
	e.GET("/:users", getreq)
	//Delete Data
	e.DELETE("/:users", deletereq)
	//Put Data
	e.PUT("/:users", putreq)
	//start server
	e.Logger.Fatal(e.Start(":8080"))
}
