package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/stevengm45/clone-twitter-go-react/db"
)

/*LeoTweetsSeguidores lee los tweets de todos nuestros seguidores*/
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Deve enviar el aprametro pagina com entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := db.LeoTweetsSeguidores(IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
