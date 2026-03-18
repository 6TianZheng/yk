package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// Logging 中间件示例
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 在这里可以加入日志
		next.ServeHTTP(w, r)
	})
}
