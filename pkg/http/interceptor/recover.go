package interceptor

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func RecoverHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				logEvent := log.Error().Caller()
				if err, ok := r.(error); ok {
					logEvent.Err(err)
				}
				logEvent.Msgf("Recover from panic %v", r)
				http.Error(w, "Произошла ошибка при обработке запроса. Попробуйте еще раз", http.StatusInternalServerError)
			}
		}()
		fn(w, r)
	}
}
