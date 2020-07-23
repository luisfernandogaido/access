package model

import (
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	maxCountMinute = 1
)

var ErrToManyRequests = errors.New("too many requests")

type ApiAccess struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Token  string             `json:"token" bson:"token"`
	Addr   string             `json:"addr" bson:"addr"`
	Date   time.Time          `json:"date" bson:"date"`
	Minute string             `json:"minute" bson:"minute"`
}

func ApiAccessInsertMinute(token, addr string) error {
	now := time.Now()
	minute := now.Format("20060102-1504")
	aa := ApiAccess{
		Id:     primitive.NewObjectID(),
		Token:  token,
		Addr:   addr,
		Date:   now,
		Minute: minute,
	}
	_, err := db.Collection("api_access").InsertOne(nil, aa)
	if err != nil {
		return fmt.Errorf("apiaccessinsertminute: %w", err)
	}
	return nil
}

func ApiAcessCountMinute(token string) (int, error) {
	now := time.Now()
	minute := now.Format("20060102-1504")
	filter := bson.M{
		"token":  token,
		"minute": minute,
	}
	count, err := db.Collection("api_access").CountDocuments(nil, filter)
	if err != nil {
		return 0, fmt.Errorf("apiacesscountminute: %w", err)
	}
	return int(count), nil
}

func ApiAcessInsertCountMinute(token, addr string) error {
	if err := ApiAccessInsertMinute(token, addr); err != nil {
		return fmt.Errorf("apiacessinsertcountminute: %w", err)
	}
	count, err := ApiAcessCountMinute(token)
	if err != nil {
		return fmt.Errorf("apiacessinsertcountminute: %w", err)
	}
	if count > maxCountMinute {
		return ErrToManyRequests
	}
	return nil
}
