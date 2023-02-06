package routes

import (
	"github.com/PedroDalpa/primeiro-crud-golang/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:id", user.FindById)
	r.GET("/user/email/:email", user.FindByEmail)
	r.POST("/user", user.Create)
	r.PUT("/user/:id", user.Update)
	r.DELETE("/user/:id", user.Delete)

}
