package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User parameters
type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FirstName string             `json:"first_name" bson:"first_name,omitempty"`
	LastName  string             `json:"last_name" bson:"last_name,omitempty"`
	Phone     string             `json:"phone,omitempty"`
	Email     string             `json:"email,omitempty"`
	Password  string             `json:"password,omitempty"`
	Address   string             `json:"address"`
	City      string             `json:"city"`
	Country   string             `json:"country"`
	Role      string             `json:"role,omitempty"`
	PicURL    string             `json:"pic_url" bson:"pic_url,omitempty"`
	Location  struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"location"`
	CreatedAt int64  `json:"created_at"`
	Paragraph string `json:"paragraph"`
	Offers    string `json:"offers"`
	LastIP    string `json:"last_ip,omitempty" bson:"last_ip,omitempty"`
	CFCookie  string `json:"cf_cookie,omitempty" bson:"cf_cookie,omitempty"`
}

// RegisterUser .
type RegisterUser struct {
	FirstName string `json:"first_name" bson:"first_name,omitempty"`
	LastName  string `json:"last_name" bson:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	PicURL    string `json:"pic_url" bson:"pic_url,omitempty"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Paragraph string `json:"paragraph"`
	Offers    string `json:"offers"`
	Location  struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	} `json:"location"`
	LastIP   string `json:"last_ip,omitempty" bson:"last_ip,omitempty"`
	CFCookie string `json:"cf_cookie,omitempty" bson:"cf_cookie,omitempty"`
}

// Loginuser .
type LoginUser struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// UserUpdate parameters
type UserUpdate struct {
	Address   string `json:"address,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	PicURL    string `json:"pic_url" bson:"pic_url,omitempty"`
	Paragraph string `json:"paragraph,omitempty"`
	Offers    string `json:"offers"`
	Location  struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	} `json:"location,omitempty"`
	LastIP   string `json:"last_ip" bson:"last_ip"`
	CFCookie string `json:"cf_cookie" bson:"cf_cookie"`
}
