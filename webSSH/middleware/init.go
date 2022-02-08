package middleware

import "github.com/gin-gonic/gin"

func InitMiddler(r *gin.Engine) {
	r.Use(Options)
}
