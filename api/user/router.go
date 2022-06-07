package user

import (
	dependency "github.com/JoaoGabrielManochio/webapi-go/config"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {

	business := NewUserBusiness(dependency.UserRepository)
	controller := NewUserController(business)

	g.POST("", controller.CreateUser)

	// user := main.Group("users")
	// {
	// 	user.GET("/:id", controllers.ShowUser)
	// 	user.GET("/", controllers.ShowUsers)
	// 	user.POST("/", controllers.CreateUser)
	// 	user.PUT("/", controllers.UpdateUser)
	// 	user.DELETE("/:id", controllers.DeleteUser)
	// }
}
