package routes

import (
	"github.com/JoaoGabrielManochio/webapi-go/api/user"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) {
	main := router.Group("api/v1")
	{
		user.Router(main.Group("user"))

		// user := main.Group("users")
		// {
		// 	user.GET("/:id", controllers.ShowUser)
		// 	user.GET("/", controllers.ShowUsers)
		// 	user.POST("/", controllers.CreateUser)
		// 	user.PUT("/", controllers.UpdateUser)
		// 	user.DELETE("/:id", controllers.DeleteUser)
		// }
	}
}
