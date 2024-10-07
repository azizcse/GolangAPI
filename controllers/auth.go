package controllers

import (
	"ginapi/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func InitAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: *authService,
	}
}

func (a *AuthController) InitRoute(router *gin.Engine) {
	routers := router.Group("/auth")
	routers.POST("/login", a.Login())
	routers.POST("/register", a.Register())
}

func (a *AuthController) Login() gin.HandlerFunc {
	type RegisterBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var registerBody RegisterBody
		if err := ctx.BindJSON(&registerBody); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		user, err := a.authService.Login(&registerBody.Email, &registerBody.Password)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": user,
		})
	}
}

func (a *AuthController) Register() gin.HandlerFunc {
	type RegisterBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var registerBody RegisterBody
		if err := ctx.BindJSON(&registerBody); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		user, err := a.authService.Register(&registerBody.Email, &registerBody.Password)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": user,
		})
	}
}
