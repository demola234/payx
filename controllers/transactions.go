package controllers

import(
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go-service/payx/interfaces"
	"go-service/payx/database"
	"go-service/payx/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

var transactionCollection *mongo.Collection = database.PayxCollection(database.Client, "Transaction")

func Deposit() gin.HandlerFunc{

	return func(c *gin.Context){

		var body interfaces.DepositPayload
		err := c.BindJSON(&body)

		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		validationErr := validate.Struct(body)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}
	
		var url = "https://api.paystack.co/transaction/initialize"

		// parse user details
		email := c.MustGet("email").(string)
		account_number := c.MustGet("account_number").(string)

		payload, _ := json.Marshal(&interfaces.Deposit{
			Amount:body.Amount*100,
			Email:email,
			Metadata: interfaces.Metadata{
				Amount: body.Amount,
				Message: body.Message,
				DebitorAccount: "0",   //payx account number
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
		if err != nil{
			panic(err)
		}
		defer resp.Body.Close()
		depositResponse := new(interfaces.DepositResponse)
		json.NewDecoder(resp.Body).Decode(depositResponse)
		c.JSON(http.StatusOK, gin.H{"data": depositResponse.Data.AuthorizationUrl})
}
}

func Verify() gin.HandlerFunc{

	return func(c *gin.Context){

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		ref := c.Param("ref")
		count, err := transactionCollection.CountDocuments(ctx, bson.M{"reference": ref})

		if err != nil {
			c.JSON(500, gin.H{"error": "Error processing the transaction"})
		}
		if count > 0{
			c.JSON(403, gin.H{"error": "This transaction has already been processed"})
		}

		url := "https://api.paystack.co/transaction/verify/"+ref
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("cache-control", "non-cache")
		req.Header.Set("Authorization", "Bearer sk_test_530cc30f2989b68e407c5f8997ee137e23ab40ef")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		depositResponse := new(interfaces.VerificationResponse)
		json.NewDecoder(resp.Body).Decode(depositResponse)
		// check if paystack verification was confirmed
		if !depositResponse.Status{
			c.JSON(403, gin.H{"data":"Invalid verification code"})
		}
		verificationErr := saveTransaction(depositResponse.Data.Metadata, ref)
		if verificationErr != nil{
			c.JSON(500, gin.H{"data":"Error occured while saving the transaction"})
			
		}
		c.JSON(http.StatusOK, gin.H{"data": "Success"})
	}
}

func saveTransaction(payload interfaces.VerificationResponseDataMetadata, reference string) (error){

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	amount, _ := strconv.Atoi(payload.Amount)
	creditor, _ := strconv.Atoi(payload.CreditorAccount)
	debitor, _ := strconv.Atoi(payload.DebitorAccount)
	
	var transaction models.Transaction

	transaction.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	transaction.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	transaction.ID = primitive.NewObjectID()
	transaction.Transaction_id = transaction.ID.Hex()
	transaction.Reference = reference
	transaction.Message = payload.Message
	transaction.Amount = amount
	transaction.Sender_Acct_Number = creditor
	transaction.Receiver_Acct_Number = debitor
	transaction.Transfer_Status = true
	_, error := transactionCollection.InsertOne(ctx, transaction)
	return error
}



func GetUserTransaction() gin.HandlerFunc{

	return func(c *gin.Context){

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		account_number := c.MustGet("account_number").(string)
		account, _ := strconv.Atoi(account_number)
		fmt.Print(account)
		matchStage := bson.D{{"$match", bson.D{{"sender_acct_number", account}, {"receiver_acct_number", account}}}}
		// matchStage1 := bson.D{{"$match", bson.D{{"receiver_acct_number", account}}}}

		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_count", bson.D{{"$sum", 1}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}

		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"Transactions", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			}}}

			sortStage := bson.D{{"$sort", bson.D{{"_id", -1}}}}

		result, err := transactionCollection.Aggregate(ctx, mongo.Pipeline{matchStage, sortStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			c.JSON(500, gin.H{"status": "Failure",
				"message": "An error Occurred while listing user transaction"})

		}
		var allUsers []bson.M
		if err = result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)
		}

		c.JSON(200, gin.H{"status": "Success", "data": allUsers[0]})

	}
}