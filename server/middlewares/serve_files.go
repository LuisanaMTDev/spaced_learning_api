package middlewares

import "net/http"

func ExcludeFiles(next http.Handler, excluded []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, path := range excluded {
			if r.URL.Path == "/"+path || r.URL.Path == path {
				http.NotFound(w, r)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
