package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID              primitive.ObjectID `bson:"_id"`
	Account_Number  string             `json:"account_number" validate:"required"`
	Account_Balance uint               `json:"account_balance"`
	Created_at      time.Time          `json:"created_at"`
	Updated_at      time.Time          `json:"updated_at"`
	Account_Id      string             `json:"account_id"`
}
