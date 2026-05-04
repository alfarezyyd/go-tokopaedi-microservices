package user

import (
	"go-tokopaedi-microservices/model"

	"github.com/gin-gonic/gin"
)

type Service interface {
	FindAll(ginContext *gin.Context) []*model.UserResponse
	Register(ginContext *gin.Context, createUserRequest *model.CreateUserRequest) []*model.UserResponse
}
