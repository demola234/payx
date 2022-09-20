package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	First_name     *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name      *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password       *string            `json:"password" validate:"required,min=6"`
	Email          *string            `json:"email" validate:"required"`
	Image          *string            `json:"image"`
	Phone          *string            `json:"phone" validate:"required"`
	Address        *string            `json:"address"`
	BVN            *string            `json:"bvn"`
	Account_Number *string            `json:"account_number"`
	Balance        *uint64               `json:"balance"`
	Token          *string            `json:"token"`
	Refresh_Token  *string            `json:"refresh_token"`
	Created_at     time.Time          `json:"created_at"`
	Updated_at     time.Time          `json:"updated_at"`
	User_id        string             `json:"user_id"`
	Card_id        *string            `json:"card_id"`
	Account_id     *string            `json:"account_id"`
}
