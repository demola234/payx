package controllers

import (
	// "context"
	// "time"
	"context"
	"go-service/payx/interfaces"
	"go-service/payx/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "net/url"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Deposit() gin.HandlerFunc {

	return func(c *gin.Context) {

		var body interfaces.DepositPayload
		err := c.BindJSON(&body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		validationErr := validate.Struct(body)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}

		var url = "https://api.paystack.co/transaction/initialize"

		// parse user details
		email := c.MustGet("email").(string)
		account_number := c.MustGet("account_number").(string)

		payload, _ := json.Marshal(&interfaces.Deposit{
			Amount: body.Amount,
			Email:  email,
			Metadata: interfaces.Metadata{
				Amount:          body.Amount,
				Message:         body.Message,
				DebitorAccount:  "0", //payx account number
				CreditorAccount: account_number,
			},
		})

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer sk_test_530cc30f2989b68e407c5f8997ee137e23ab40ef")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		depositResponse := new(interfaces.DepositResponse)
		json.NewDecoder(resp.Body).Decode(depositResponse)
		c.JSON(http.StatusOK, gin.H{"data": depositResponse.Data.AuthorizationUrl})
	}
}

// Ademola
func GetDeposit() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// Bolu
func WithdrawFunds() {}

// Ademola
func TransferFunds() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		debitNumber := c.PostForm("debit_number")
		creditorNumber := c.PostForm("credit_number")
		amount := c.PostForm("amount")
		// message := c.PostForm("message")

		am, err := strconv.ParseUint(amount, 10, 32)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error converting amount to unit"})
		}

		var debitAccount models.Account
		var creditorAccount models.Account

		var updateObj primitive.D
		var updateObj1 primitive.D

		err1 := accountCollection.FindOne(ctx, bson.M{"account_number": debitNumber}).Decode(&debitAccount)

		defer cancel()
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
		}

		err2 := accountCollection.FindOne(ctx, bson.M{"account_number": creditorNumber}).Decode(&creditorAccount)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
		}

		if debitAccount.Account_Balance >= am {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Insufficient Funds"})
		}
		filter1 := bson.M{"account_number": debitNumber}

		debit := debitAccount.Account_Balance - am
		if creditorAccount.Account_Balance != 0 {
			updateObj = append(updateObj, bson.E{"account_balance", debit})
		}

		debitAccount.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", debitAccount.Updated_at})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		_, updateErr1 := accountCollection.UpdateOne(
			ctx,
			filter1,
			bson.D{{"$set", updateObj}},
			&opt,
		)

		if updateErr1 != nil {
			// msg := "User Failed"
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		filter2 := bson.M{"account_number": creditorNumber}

		creditor := creditorAccount.Account_Balance + am
		if creditorAccount.Account_Balance != 0 {
			updateObj1 = append(updateObj, bson.E{"account_balance", creditor})
		}

		creditorAccount.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj1 = append(updateObj, bson.E{"updated_at", creditorAccount.Updated_at})

		opt = options.UpdateOptions{
			Upsert: &upsert,
		}

		_, updateErr2 := accountCollection.UpdateOne(
			ctx,
			filter2,
			bson.D{{"$set", updateObj1}},
			&opt,
		)

		if updateErr2 != nil {
			// msg := "User Failed"
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		//Add to Transaction History

		c.JSON(http.StatusOK, gin.H{"data": "successfully transferred!!!"})

	}
}

// Ademola
func AirTime() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Bolu
func UtilityBills() {}
