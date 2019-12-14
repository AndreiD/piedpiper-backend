package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"piedpiper/models"
	"piedpiper/utils/log"
	"time"
)

// ListUsersDB ..
func ListUsersDB(page int, limit int) ([]models.User, error) {
	start := time.Now()
	var users []models.User
	collection := db.Collection("users")

	filter := bson.M{"role": "user"}

	opts := &options.FindOptions{
		Projection: bson.M{"password": 0, "last_ip": 0, "cf_cookie": 0, "role": 0},
	}
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64((page - 1) * limit))
	opts.SetSort(bson.M{"created_at": -1}) // show new records on top

	cur, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return users, err
	}
	for cur.Next(context.Background()) {
		var user models.User
		err = cur.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	log.Printf("listing all users took %s", time.Since(start))
	return users, nil
}
