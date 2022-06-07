package routes

import (
	"github.com/JoaoGabrielManochio/webapi-go/api/user"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) {
	main := router.Group("api/v1")
	{
		user.Router(main.Group("user"))
	}
}
