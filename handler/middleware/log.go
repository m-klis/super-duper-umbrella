package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func LoggerHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()
		defer func() {
			since := time.Since(start).Milliseconds()
			status := ww.Status()

			log.Printf("request completed in %dms: HTTP %d", since, status)
		}()
		next.ServeHTTP(ww, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
