package handlers

import (
	"github.com/Clever/sphinx/ratelimiter"
	"net/http"
)

// START OMIT
func NewHTTPLimiter(rateLimiter ratelimiter.RateLimiter, proxy http.Handler) http.Handler {
	return &httpRateLimiter{rateLimiter: rateLimiter, proxy: proxy}
}

func NewHTTPLogger(rateLimiter ratelimiter.RateLimiter, proxy http.Handler) http.Handler {
	return &httpRateLogger{rateLimiter: rateLimiter, proxy: proxy}
}

// END OMIT
