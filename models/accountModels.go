package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID              primitive.ObjectID `bson:"_id"`
	Account_Number  *int               `json:"account_number" validate:"required"`
	Account_Balance *int               `json:"account_balance"`
	Created_at      time.Time          `json:"created_at"`
	Updated_at      time.Time          `json:"updated_at"`
	Transaction_ID  string             `json:"transaction_id"`
}
