package user

import (
	"context"
	"fmt"
	"go-tokopaedi-microservices/configs"
	"go-tokopaedi-microservices/model"
	"go-tokopaedi-microservices/pkg/exception"
	"go-tokopaedi-microservices/pkg/helper"
	"go-tokopaedi-microservices/protobuf/genproto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	dbConnection       *gorm.DB
	viperConfig        *viper.Viper
	redisInstance      *configs.RedisInstance
	grpcClientRegistry *configs.GrpcClientRegistry
}

func NewService(dbConnection *gorm.DB,
	viperConfig *viper.Viper,
	redisInstance *configs.RedisInstance,
	grpcClientRegistry *configs.GrpcClientRegistry,
) *ServiceImpl {
	return &ServiceImpl{
		dbConnection:       dbConnection,
		viperConfig:        viperConfig,
		redisInstance:      redisInstance,
		grpcClientRegistry: grpcClientRegistry,
	}
}

func (userService *ServiceImpl) FindAll(ginContext *gin.Context) []*model.UserResponse {
	var allUser []*model.UserResponse
	err := userService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		grpcClient, err := userService.grpcClientRegistry.Get("user")
		assertGrpcClient := grpcClient.(genproto.UserServiceClient)
		contextWithTimeout, cancelFunc := context.WithTimeout(ginContext.Request.Context(), 10*time.Second)
		defer cancelFunc()
		userResponses, err := assertGrpcClient.FindAll(contextWithTimeout, &emptypb.Empty{})
		helper.CheckErrorOperation(err, exception.NewApplicationError(500, exception.MsgInternalError))
		fmt.Println(userResponses)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return allUser
}

func (userService *ServiceImpl) Register(ginContext *gin.Context, createUserRequest *model.CreateUserRequest) []*model.UserResponse {
	var allUser []*model.UserResponse
	err := userService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		grpcClient, err := userService.grpcClientRegistry.Get("user")
		assertGrpcClient := grpcClient.(genproto.UserServiceClient)
		fmt.Println(grpcClient, assertGrpcClient, err)
		helper.CheckErrorOperation(err, exception.NewApplicationError(http.StatusInternalServerError, exception.MsgInternalError))
		contextWithTimeout, cancel := context.WithTimeout(ginContext, 10*time.Second)
		defer cancel()
		grpcCreateUserRequest := helper.MapCreateRequestIntoEntity[model.CreateUserRequest, genproto.CreateUserRequest](createUserRequest)
		registerResult, err := assertGrpcClient.HandleRegister(contextWithTimeout, grpcCreateUserRequest)
		fmt.Println(registerResult, err)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return allUser
}
