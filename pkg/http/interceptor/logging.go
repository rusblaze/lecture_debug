package interceptor

import (
	"fmt"
	"lecture/pkg/log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func LogHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError) // Отдаем 500 клиенту
			return
		}
		log.Trace(r.Context()).Msgf("Request is %q", x) // Логируем запрос с уровнем TRACE
		rec := httptest.NewRecorder()
		fn(rec, r)
		log.Trace(r.Context()).Msgf("Response is %q", rec.Body) // Логируем ответ с уровнем TRACE

		// this copies the recorded response to the response writer
		for k, v := range rec.Result().Header {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		rec.Body.WriteTo(w)
	}
}
