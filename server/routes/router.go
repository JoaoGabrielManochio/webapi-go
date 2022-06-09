package routes

import (
	"github.com/JoaoGabrielManochio/webapi-go/api/transaction"
	"github.com/JoaoGabrielManochio/webapi-go/api/user"
	"github.com/JoaoGabrielManochio/webapi-go/api/wallet"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) {
	main := router.Group("api/v1")
	{
		user.Router(main.Group("user"))
		wallet.Router(main.Group("wallet"))
		transaction.Router(main.Group("transaction"))
	}
}
