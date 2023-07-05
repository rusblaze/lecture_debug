package interceptor

import (
	"lecture/pkg/log"
	"net/http"
)

func RecoverHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				logEvent := log.Error(r.Context()).Caller()
				if err, ok := rec.(error); ok {
					logEvent.Err(err)
				}
				logEvent.Msgf("Recover from panic %v", rec)
				http.Error(w, "Произошла ошибка при обработке запроса. Попробуйте еще раз", http.StatusInternalServerError)
			}
		}()
		fn(w, r)
	}
}
