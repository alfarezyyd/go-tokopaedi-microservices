package routes

import (
	"go-tokopaedi-microservices/services/gateway/user"

	"github.com/gin-gonic/gin"
)

type AuthenticationRoutes struct {
	userHandler user.Handler
}

func NewAuthenticationRoutes(userHandler user.Handler) *AuthenticationRoutes {
	return &AuthenticationRoutes{
		userHandler: userHandler,
	}
}

func (authenticationRoutes *AuthenticationRoutes) Setup(routerGroup *gin.RouterGroup) {
	userRouterGroup := routerGroup.Group("users")
	userRouterGroup.GET("", authenticationRoutes.userHandler.FindAll)
	userRouterGroup.POST("", authenticationRoutes.userHandler.HandleRegister)
}
