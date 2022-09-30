package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/stevengm45/clone-twitter-go-react/db"
	"github.com/stevengm45/clone-twitter-go-react/models"
)

/*SubirAvatar sube el Avatar al servidor*/
func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.User
	var status bool

	usuario.Avatar = IDUsuario + "." + extension
	status, err = db.ModificoRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al grabar el avatar en la DB! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
