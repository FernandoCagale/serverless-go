package middleware

import (
	"net/http"
	"os"

	"github.com/FernandoCagale/serverless-go/src/datastore"
	"github.com/gorilla/context"
)

func BindDb(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := datastore.New(os.Getenv("DATASTORE_URL"))
		if err != nil {
			panic(err)
		}

		context.Set(r, "db", db)

		defer db.Close()

		next.ServeHTTP(w, r)
	})
}
