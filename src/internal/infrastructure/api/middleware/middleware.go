package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/tayusa/notugly_backend/configs"
	"github.com/tayusa/notugly_backend/internal/infrastructure/api/property"
	"github.com/tayusa/notugly_backend/pkg/firebase"
)

func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token, err := firebase.FetchToken(r)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error verifying ID token\n"))
			return
		}

		next(
			w,
			r.WithContext(
				property.SetUserId(r.Context(), token.UID)),
			p)
	}
}

func SetHeader(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		origin := fmt.Sprintf(
			"http://%s:%s", configs.Data.Frontend.Host, configs.Data.Frontend.Port)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		next(w, r, p)
	}
}
