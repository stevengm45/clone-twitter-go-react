package db

import (
	"context"
	"fmt"
	"time"

	"github.com/stevengm45/clone-twitter-go-react/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultaRelacion consulta la relacion entre 2 usuarios*/
func COnsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionud": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}
	return true, nil

}
