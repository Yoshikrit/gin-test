package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func Cors() gin.HandlerFunc {
	return cors.Default()
}