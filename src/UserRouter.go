package src

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/golangAPI-gin/service"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	user.GET("/", service.FindAllUsers)
	user.GET("/:id", service.FindUserById)
	user.POST("/", service.PostUser)
	user.DELETE("/:id", service.DeleteUser)
	user.PUT("/:id", service.PutUser)
}


