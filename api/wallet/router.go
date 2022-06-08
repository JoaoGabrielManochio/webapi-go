package wallet

import (
	dependency "github.com/JoaoGabrielManochio/webapi-go/config"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup) {

	business := NewWalletBusiness(dependency.WalletRepository)
	controller := NewWalletController(business)

	g.POST("", controller.CreateWallet)
	g.GET("/:id", controller.GetWallet)
	g.GET("", controller.GetWallets)
	g.PUT("/", controller.UpdateWallet)
	g.DELETE("/:id", controller.DeleteWallet)
}
