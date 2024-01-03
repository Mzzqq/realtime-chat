package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/user"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
}
