package transaction

import (
	dependency "github.com/JoaoGabrielManochio/webapi-go/config"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {

	business := NewTransactionBusiness(dependency.TransactionRepository, dependency.TransactionService)
	controller := NewTransactionController(business)

	g.POST("", controller.CreateTransaction)
	g.GET("/:id", controller.GetTransaction)
	g.GET("", controller.GetTransactions)
}
