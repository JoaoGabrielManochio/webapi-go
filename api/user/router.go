package user

import (
	dependency "github.com/JoaoGabrielManochio/webapi-go/config"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {

	business := NewUserBusiness(dependency.UserRepository)
	controller := NewUserController(business)

	g.POST("", controller.CreateUser)
	g.GET("/:id", controller.GetUser)
	g.GET("", controller.GetUsers)
	g.PUT("/", controller.UpdateUser)
	g.DELETE("/:id", controller.DeleteUser)
}
