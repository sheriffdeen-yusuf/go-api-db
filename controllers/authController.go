package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
	Store(ctx *gin.Context)
}

type authController struct{}

func AuthControllerImpl() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "we are login",
	})
}

func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "we are register",
	})
}

func (c *authController) ResetPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "we are resetpassword",
	})
}

func (c *authController) Store(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "we are resetpassword",
	})
}
