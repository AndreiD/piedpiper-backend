package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"piedpiper/models"
	"piedpiper/utils/log"
	"time"
)

// Login ..
func Login(email string, password string) (models.User, error) {
	var user models.User
	collection := db.Collection("users")

	filter := bson.M{"email": email, "password": password}

	opts := &options.FindOneOptions{
		Projection: bson.M{"email": 1, "password": 1, "role": 1, "security": 1, "first_name": 1, "last_name": 1},
	}

	time.Sleep(500 * time.Millisecond)

	err := collection.FindOne(context.Background(), filter, opts).Decode(&user)
	if err != nil {
		return user, err
	}

	log.Printf("%s %s logged in", user.FirstName, user.LastName)

	return user, nil
}
