package db

import (
	"context"
	"time"

	"github.com/stevengm45/clone-twitter-go-react/models"
)

/*InsertoRelacion graba la relación en la BD */
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
