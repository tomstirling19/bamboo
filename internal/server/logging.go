package server

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = io.ReadAll(r.Body)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		log.Printf("Received %s request for %s", r.Method, r.URL.Path)

		if len(bodyBytes) > 0 {
			log.Printf("Request body: %s", string(bodyBytes))
		} else {
			log.Println("Request has no body.")
		}

		next.ServeHTTP(w, r)
	})
}
