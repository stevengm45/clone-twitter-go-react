package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BorroTweet borra un twet determinado */
func BorroTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}
