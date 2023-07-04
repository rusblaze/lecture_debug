package interceptor

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func LogRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		fmt.Printf("%q", dump)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
