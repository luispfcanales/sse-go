package middleware

import (
	"net/http"
)

//StateUser check auth user
func StateUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
