package db

import (
	"context"
	"time"

	"github.com/stevengm45/clone-twitter-go-react/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoRegistro es la prada final con la BD ára insertat los datos del usuario */
func InsertoRegistro(u models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
