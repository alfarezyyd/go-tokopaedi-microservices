package routes

import (
	"go-tokopaedi-microservices/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ProtectedRoutes struct {
	viperConfig *viper.Viper
}

func NewProtectedRoutes(
	viperConfig *viper.Viper,

) *ProtectedRoutes {
	return &ProtectedRoutes{
		viperConfig: viperConfig,
	}
}

func (protectedRoutes *ProtectedRoutes) Setup(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middleware.AuthMiddleware(protectedRoutes.viperConfig))
}
