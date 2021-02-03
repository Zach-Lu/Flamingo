package collections

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mimigo/config"
	"time"
)

type PersonalAccessToken struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	Token     string             `json:"token" bson:"token"`
	Device    string             `json:"device" bson:"device"`
	IpAddress string             `json:"ipAddress" bson:"ipAddress"`
	UserAgent string             `json:"userAgent" bson:"userAgent"`
	ExpiredAt time.Time          `json:"expired_at" bson:"expired_at"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (token PersonalAccessToken) FormatToken() PersonalAccessToken {
	token.Id = primitive.NewObjectID()
	token.ExpiredAt = time.Now().Add(config.TokenLife)
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	return token
}
