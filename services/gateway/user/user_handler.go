package user

import "github.com/gin-gonic/gin"

type Handler interface {
	FindAll(*gin.Context)
	HandleRegister(*gin.Context)
}
