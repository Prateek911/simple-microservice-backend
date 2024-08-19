package api

import (
	"context"
	"net/http"
	"simple-microservice-backend/config"
	"simple-microservice-backend/db"
	"time"
)

type contextKey string

const (
	dbContextKey contextKey = "db"
)

func ApplicationContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(string(ACCESS_CONTROL_ALLOW_ORIGIN), "*")
		w.Header().Add(string(ACCESS_CONTROL_ALLOW_CREDENTIALS), "true")
		w.Header().Add(string(ACCESS_CONTROL_ALLOW_HEADERS), string(DEFAULT_HEADERS))
		w.Header().Add(string(ACCESS_CONTROL_ALLOW_METHODS), "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		opts, err := config.NewServerConfig()
		if err != nil {
			http.Error(w, "Error initialising API Handler ", http.StatusInternalServerError)
			return
		}

		timeOutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(opts.ContextTimeOut)*time.Second)
		defer cancel()
		ctx := context.WithValue(r.Context(), dbContextKey, db.DB.WithContext(timeOutCtx))
		next.ServeHTTP(w, r.WithContext(ctx))
		//post-processing
	})
}
