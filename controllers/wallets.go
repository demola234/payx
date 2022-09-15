package controllers

import (
	"context"
	"go-service/payx/database"
	"go-service/payx/models"
	"go-service/payx/utils"

	// "net/http"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var accountCollection *mongo.Collection = database.PayxCollection(database.Client, "Account")
var cardCollection *mongo.Collection = database.PayxCollection(database.Client, "Card")

func CreateAccountDetails(c *gin.Context) (models.Account, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var account models.Account
	err := c.BindJSON(&account)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": err.Error(),
		// })

	}
	error := validate.Struct(account)

	if error != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}

	account.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	account.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	account.Account_Balance = 0.0
	account.ID = primitive.NewObjectID()
	account.Account_Id = account.ID.Hex()
	accountNumber := utils.GenerateRandomString(10, 2)
	account.Account_Number = accountNumber

	_, error = accountCollection.InsertOne(ctx, account)
	return account, error
}

func CreateUsersCard(c *gin.Context) (models.Card, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var card models.Card
	_ = c.BindJSON(&card)

	_ = validate.Struct(card)

	card.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	card.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	card.Card_Type = "VISA"
	card.ID = primitive.NewObjectID()
	card.Card_ID = card.ID.Hex()
	cardNumber := utils.GenerateRandomString(16, 2)
	card.Card_Number = cardNumber

	_, _ = cardCollection.InsertOne(ctx, card)
	return card, nil
}

func GetUserAccountDetailsByID() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		account_Id := c.Param("account_id")

		var foundAccount models.Account

		err := accountCollection.FindOne(ctx, bson.M{"account_id": account_Id}).Decode(&foundAccount)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
		}
		c.JSON(http.StatusOK, foundAccount)
	}
}


func GetUserAccountDetailsByNumber() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		account_Number := c.Param("account_number")

		var foundAccount models.Account

		err := accountCollection.FindOne(ctx, bson.M{"account_number": account_Number}).Decode(&foundAccount)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
		}
		c.JSON(http.StatusOK, foundAccount)
	}
}

func GetUserCardDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		card_id := c.Param("card_id")

		var foundCard models.Card

		err := cardCollection.FindOne(ctx, bson.M{"card_id": card_id}).Decode(&foundCard)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
		}
		c.JSON(http.StatusOK, foundCard)
	}
}
