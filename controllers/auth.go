package controllers

import "github.com/gin-gonic/gin"

type AuthController struct {
}

func InitAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) InitRoute(router *gin.Engine) {
	routers := router.Group("/auth")
	routers.POST("/login", a.Nope())
	routers.POST("/register")
}

func (*AuthController) Nope() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Connected",
		})
	}
}
