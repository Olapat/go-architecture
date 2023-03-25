package tokenResetPasswordEntity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenResetPassword struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	ExpiresAt time.Time          `json:"expires_at" bson:"expires_at"`
	Token     string             `json:"token" bson:"token"`
	IsConfirm bool               `json:"is_confirm" bson:"is_confirm"`
}
