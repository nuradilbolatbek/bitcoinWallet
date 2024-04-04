package handler

import (
	"bitcoinWallet/package/service"
	"github.com/gin-gonic/gin"
	"sync"
	//"path/to/your/project/service" // Replace with your actual service package
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
	mu       sync.Mutex
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		wallets := api.Group("/wallets")
		{
			wallets.POST("/", h.createWallet)
			wallets.GET("/:id", h.getWalletByUserId)
			wallets.PUT("/:id/deposit", h.depositToWallet)
			wallets.PUT("/:id/withdraw", h.withdrawFromWallet)
		}
	}

	return router
}
