package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID          primitive.ObjectID `bson:"_id"`
	Card_Number *int               `json:"card_number"`
	Card_Type   *int               `json:"card_type" validate:"required,eq=VISA|eq=MASTER"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
	Card_ID     string             `json:"card_id"`
}
