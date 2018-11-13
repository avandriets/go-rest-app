package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
	User        string
	AccessToken string
	Secure      bool
}

func (amw *AuthMiddleware) SetSecureOn() {
	amw.Secure = true
}

func (amw *AuthMiddleware) SetSecureOff() {
	amw.Secure = false
}

func (amw *AuthMiddleware) CheckJWTToken(token string) bool {
	if !amw.Secure {
		return true
	}

	// TODO implement checking JWT token
	return false
}

func (amw *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authentication")

		if authenticated := amw.CheckJWTToken(token); authenticated {
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
