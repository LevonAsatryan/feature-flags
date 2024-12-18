package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Keeping this just in case I find better solution
// func ValidateRequestBody(ctx *gin.Context) {
// 	var body models.Group
// 	if err := ctx.ShouldBindJSON(&body); err != nil {
// 		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.Set("validatedBody", body)
// 	ctx.Next()
// }

func ValidateId(ctx *gin.Context) {
	if _, err := uuid.Parse(ctx.Param("id")); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Next()
}
