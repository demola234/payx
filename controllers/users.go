package controllers

import (
	"context"
	"fmt"
	"go-service/payx/database"
	"go-service/payx/helpers"
	"go-service/payx/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()
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
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var users models.User
		//convert the JSON data coming from postman to something that golang understands
		err := c.BindJSON(&users)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		//validate the data based on user struct
		validationErr := validate.Struct(users)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		//you'll check if the email has already been used by another user
		count1, err := userCollection.CountDocuments(ctx, bson.M{"email": users.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count1 > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email  already exits"})
			return
		}

		//you'll also check if the phone no. has already been used by another user
		count2, err := userCollection.CountDocuments(ctx, bson.M{"phone": users.Phone})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count2 > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this phone number  already exits"})
			return
		}

		//hash password
		password := HashPassword(*users.Password)
		users.Password = &password

		//create some extra details for the user object - created_at, updated_at, ID
		users.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		users.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		users.ID = primitive.NewObjectID()
		users.User_id = users.ID.Hex()
		//generate token and refresh token (generate all tokens function from helper)
		token, refreshToken, _ := helpers.GenerateAllTokens(*users.Email, *users.First_name, *users.Last_name, users.User_id)
		users.Token = &token
		users.Refresh_Token = &refreshToken

		//Create User Account
		accountDetails, error := CreateAccountDetails(c)
		if error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "Could not create Account number",
			})
		}
		users.Account_Number = &accountDetails.Account_Number
		users.Account_id = &accountDetails.Account_Id
		users.Balance = &accountDetails.Account_Balance

		cardDetails, error := CreateUsersCard(c)
		if error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "Could not create Account number",
			})
		}

		users.Card_id = &cardDetails.Card_ID

		//if all ok, then you insert this new user into the user collection
		_, err = userCollection.InsertOne(ctx, users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		//return status OK and send the result back
		c.IndentedJSON(http.StatusCreated, users)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// defer cancel()
		//convert the login data from postman which is in JSON to golang readable format
		//find a user with that email and see if that user even exists
		//then you will verify the password
		//if all goes well, then you'll generate tokens
		//update tokens - token and refresh token
		//return statusOK
	}
}

func UploadProfileImage() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {
	// Hashing the users password with bcrypt
	bytes, error := bcrypt.GenerateFromPassword([]byte(password), 15)
	if error != nil {
		log.Panic(error)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (string, bool) {
	// Compare Users Password and Provided Password
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or password is incorrect")
		check = false
	}
	return msg, check
}
