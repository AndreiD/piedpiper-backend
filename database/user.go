package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"piedpiper/models"
	"piedpiper/utils/log"
	"strconv"
	"time"
)

func CreateUser(userRegister models.RegisterUser) error {
	var user models.User

	lat, err := strconv.ParseFloat(userRegister.Location.Lat, 64)
	if err != nil {
		return err
	}
	lng, err := strconv.ParseFloat(userRegister.Location.Lng, 64)
	if err != nil {
		return err
	}

	user.FirstName = userRegister.FirstName
	user.LastName = userRegister.LastName
	user.Address = userRegister.Address
	user.City = userRegister.City
	user.Country = userRegister.Country
	user.Email = userRegister.Email
	user.Phone = userRegister.Phone
	user.Password = userRegister.Password
	user.Paragraph = userRegister.Paragraph
	user.Offers = userRegister.Offers
	user.CreatedAt = time.Now().Unix()
	user.Location.Lat = lat
	user.Location.Lng = lng
	user.Role = "user"
	user.PicURL = userRegister.PicURL
	user.LastIP = userRegister.LastIP
	user.CFCookie = userRegister.CFCookie

	collection := db.Collection("users")

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	log.Printf("user created with ID %s", result.InsertedID)
	return nil
}

// GetUser ..
func GetUser(userID string) (models.User, error) {
	var user models.User
	collection := db.Collection("users")

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return user, err
	}
	filter := bson.M{"_id": objID}

	opts := &options.FindOneOptions{
		Projection: bson.M{"password": 0, "security": 0, "email": 0, "role": 0, "notes": 0, "last_ip": 0, "cf_cookie": 0},
	}

	err = collection.FindOne(context.Background(), filter, opts).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetAuthenticatedUser ..
func GetAuthenticatedUser(userID string) (models.User, error) {
	var user models.User
	collection := db.Collection("users")

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return user, err
	}
	filter := bson.M{"_id": objID}

	opts := &options.FindOneOptions{
		Projection: bson.M{"password": 0, "role": 0, "last_ip": 0, "cf_cookie": 0},
	}

	err = collection.FindOne(context.Background(), filter, opts).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUser ..
func UpdateUser(userID string, payload models.UserUpdate) error {
	collection := db.Collection("users")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	lat, err := strconv.ParseFloat(payload.Location.Lat, 64)
	if err != nil {
		return err
	}
	lng, err := strconv.ParseFloat(payload.Location.Lng, 64)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"address":   payload.Address,
			"city":      payload.City,
			"country":   payload.Country,
			"pic_url":   payload.PicURL,
			"paragraph": payload.Paragraph,
			"offers":    payload.Paragraph,
			"location": bson.M{
				"lat": lat,
				"lng": lng,
			},
			"last_ip":   payload.LastIP,
			"cf_cookie": payload.CFCookie,
		},
	}
	result, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return err
	}
	log.Printf("%s updated. modified count %d", userID, result.ModifiedCount)
	return nil
}

// UpdateOneField ..
func UpdateOneField(userID string, update bson.M) error {
	collection := db.Collection("users")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	result, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return err
	}
	log.Printf("%s updated. modified count %d", userID, result.ModifiedCount)
	return nil
}
