package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"piedpiper/models"
	"piedpiper/utils/log"
	"time"
)

// CreateChat
func CreateChat(chat models.Chat) error {
	collection := db.Collection("chat")

	result, err := collection.InsertOne(context.Background(), chat)
	if err != nil {
		return err
	}
	log.Printf("chat created with ID %s", result.InsertedID)
	return nil
}

// GetUser ..
func ListUsersChats(uID string, page int, limit int) ([]models.Chat, error) {
	start := time.Now()
	var chatList []models.Chat
	collection := db.Collection("chat")

	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"from_user_id", uID}},
			bson.D{{"to_user_id", uID}},
		}},
	}

	opts := &options.FindOptions{
		Projection: bson.M{"nothing": 0},
	}
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64((page - 1) * limit))
	opts.SetSort(bson.M{"created_at": -1}) // show new records on top

	cur, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return chatList, err
	}
	for cur.Next(context.Background()) {
		var chat models.Chat
		err = cur.Decode(&chat)
		if err != nil {
			return chatList, err
		}
		chatList = append(chatList, chat)
	}
	log.Printf("listing all chat in this room took %s", time.Since(start))
	return chatList, nil
}
