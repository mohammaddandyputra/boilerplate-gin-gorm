package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// [ START ] HTTP 200
func ResponseOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "success", "data": data})
}

func ResponseCreated(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "success", "data": data})
}

// [ END ] HTTP 200

// [ START ] HTTP 400
func ResponseBadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": message})
}

func ResponseUnauthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
}

// [ END ] HTTP 200

// [ START ] HTTP 500
func ResponseInternalServerError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
}

// [ END ] HTTP 200
