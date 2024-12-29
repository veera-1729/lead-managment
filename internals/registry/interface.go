package registry

import "github.com/gin-gonic/gin"

type IService interface {
	ServiceName() string
	Handle(c *gin.Context)
}
