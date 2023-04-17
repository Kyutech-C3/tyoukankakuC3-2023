package routers

import (
	"github.com/Simo-C3/tyoukankakuC3-2023/api/controllers"

	"github.com/gin-gonic/gin"
)

func InitToDo(r *gin.RouterGroup) {
	u := r.Group("/todo")
	{
		u.GET("/", controllers.GetToDo)
		u.POST("/", controllers.PostToDo)
		u.PATCH("/:id", controllers.PutToDo)
		u.DELETE("/:id", controllers.DeleteToDo)
	}
}
