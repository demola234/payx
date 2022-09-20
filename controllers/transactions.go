package controllers

import (
	// "context"
	// "time"
	"go-service/payx/interfaces"

	"github.com/gin-gonic/gin"

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
func GetDeposit() {}

// Bolu
func WithdrawFunds() {}

// Ademola
func TransferFunds() {}

// Ademola
func AirTime() {}

// Bolu
func UtilityBills() {}