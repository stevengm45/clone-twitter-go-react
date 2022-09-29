package middlew

import (
	"net/http"

	"github.com/stevengm45/clone-twitter-go-react/routers"
)

/* ValidoJWT permite validar el JWT que nos viene en la peticion*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Athorization"))
		if err != nil {
			http.Error(w, "Error en el Token! "+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	}
}
