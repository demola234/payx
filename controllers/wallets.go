package controllers

import (
	"context"
	"go-service/payx/database"
	"go-service/payx/models"
	"go-service/payx/utils"

	// "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var accountCollection *mongo.Collection = database.PayxCollection(database.Client, "Account")

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

	_, error = userCollection.InsertOne(ctx, account)
	return account, error
}

func CreateUsersCard(c *gin.Context) (models.Card, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var card models.Card
	err := c.BindJSON(&card)
	if err != nil {
		// c.JSON(http.StatusBadRequest,
		// })

	}
	error := validate.Struct(card)

	if error != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}

	card.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	card.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	card.Card_Type = "VISA"
	card.ID = primitive.NewObjectID()
	card.Card_ID = card.ID.Hex()
	cardNumber := utils.GenerateRandomString(16, 2)
	card.Card_Number = cardNumber

	_, error = userCollection.InsertOne(ctx, card)
	return card, error
}
