package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TimeLogging(l *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)
		method, path, _, _ := extractInfo(ctx)
		l.Info("Request",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int64("duration", duration.Milliseconds()),
		)
	}
}

func extractInfo(ctx *gin.Context) (method string, path string, params string, query string) {
	fullPath := strings.Split(ctx.FullPath(), "/")
	return ctx.Request.Method, ctx.Request.URL.Path, fullPath[len(fullPath)-1], "" // implement query string get
}
