package middleware

import (
	"strings"
	"time"

	"github.com/DavidEsdrs/go-mercado/pkg/logger"
	"github.com/gin-gonic/gin"
)

func TimeLogging(l *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)
		method, path, _, _ := extractInfo(ctx)
		l.Info("%v %v - duration: %vms", method, path, duration.Milliseconds())
	}
}

func extractInfo(ctx *gin.Context) (method string, path string, params string, query string) {
	fullPath := strings.Split(ctx.FullPath(), "/")
	return ctx.Request.Method, ctx.Request.URL.Path, fullPath[len(fullPath)-1], "" // implement query string get
}
