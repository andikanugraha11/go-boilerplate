package controller

import "github.com/gin-gonic/gin"

type OrderController interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
}
