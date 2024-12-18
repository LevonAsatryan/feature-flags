package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: make better error handler middleware if needed
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}()

		ctx.Next()
	}
}
