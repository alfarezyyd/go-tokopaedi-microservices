package user

import (
	"go-tokopaedi-microservices/model"
	"go-tokopaedi-microservices/pkg/exception"
	"go-tokopaedi-microservices/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HandlerImpl struct {
	userService Service
	viperConfig *viper.Viper
}

func NewHandler(userService Service, viperConfig *viper.Viper) *HandlerImpl {
	return &HandlerImpl{
		userService: userService,
		viperConfig: viperConfig,
	}
}

func (userHandler *HandlerImpl) FindAll(ginContext *gin.Context) {
	all := userHandler.userService.FindAll(ginContext)
	ginContext.JSON(200, all)
}

func (userHandler *HandlerImpl) HandleRegister(ginContext *gin.Context) {
	var createUserRequest *model.CreateUserRequest
	err := ginContext.ShouldBindBodyWithJSON(&createUserRequest)
	helper.CheckErrorOperation(err, exception.NewApplicationError(http.StatusBadRequest, exception.ErrBadRequest))
	allUser := userHandler.userService.Register(ginContext, createUserRequest)
	helper.NewSuccessResponseWithEntries("User created successfully", allUser)
}
