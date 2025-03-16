package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type MiddlewareLoging struct {
	Log *logrus.Logger
}

func NewMiddlewareLoging(log *logrus.Logger) *MiddlewareLoging {
	return &MiddlewareLoging{
		Log: log,
	}
}

func (log *MiddlewareLoging) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WarpperWritter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		log.Log.Println(wrapper.StatusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
