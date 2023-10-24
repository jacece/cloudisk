package handler

import "net/http"

func AuthenInterceptor(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			username := r.Form.Get("username")
			token := r.Form.Get("token")

			if len(username) < 3 || !validToken(token) {
				http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
				return
			}
			next.ServeHTTP(w, r)
		})
}

func validToken(token string) bool {
	return len(token) == 40
}
