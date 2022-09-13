package controllers

import (
	"fmt"
	"strconv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go-service/payx/database"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var userCollection *mongo.Collection = database.PayxCollection(database.Client, "Users")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage <1{
			recordPerPage = 10
		}

		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1{
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			
		}}}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{matchStage, projectStage})
		defer cancel()
		if err != nil{
			c.JSON(500, gin.H{"status": "Failure",
								"message": "An error occured while listing user items"})
		}
		var allUsers []bson.M
		if err = result.All(ctx, &allUsers); err != nil{
			log.Fatal(err)
		}
		c.JSON(200, gin.H{"status": "Success", "data": allUsers})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var v = CreateAccountDetails()
		fmt.Sprintln(v)
	}
}

func UploadProfileImage() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func HashPassword() {}

func VerifyPassword() {

}
