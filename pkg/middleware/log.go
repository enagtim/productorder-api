package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func LogInit() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
func LoggingResultRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := &WrapperWritter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		logrus.WithFields(logrus.Fields{
			"code":   wrapper.StatusCode,
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("http api-logs")
	})
}
