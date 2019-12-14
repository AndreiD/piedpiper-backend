package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Chat parameters
type Chat struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FromUserID string             `json:"from_user_id" bson:"from_user_id,omitempty"`
	ToUserID   string             `json:"to_user_id" bson:"to_user_id,omitempty"`
	Message    string             `json:"message" bson:"message,omitempty"`
	Metadata   string             `json:"metadata" bson:"metadata,omitempty"`
	CreatedAt  int64              `json:"created_at" bson:"created_at,omitempty"`
}
