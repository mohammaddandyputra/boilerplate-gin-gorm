package controllers

import (
	"learn-gin-gorm/dto"
	"learn-gin-gorm/services"
	"learn-gin-gorm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
	userService *services.UserService
}

func NewAuthController(userService *services.UserService, authService *services.AuthService) *AuthController {
	return &AuthController{
		userService: userService,
		authService: authService,
	}
}

func (c *AuthController) ProfileUser(ctx *gin.Context) {
	email, exists := ctx.Get("email")
	if !exists {
		utils.ResponseBadRequest(ctx, "Email not found in token")
		return
	}

	user, err := c.authService.ProfileUser(email.(string))
	if err != nil {
		utils.ResponseInternalServerError(ctx, err.Error())
		return
	}

	responseData := map[string]interface{}{
		"user": user,
	}

	utils.ResponseOK(ctx, responseData)
}

func (c *AuthController) RegisterUser(ctx *gin.Context) {
	var requestBody dto.RegisterDTO

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		utils.ResponseBadRequest(ctx, err.Error())
		return
	}

	user, _ := c.userService.GetUserByEmail(requestBody.Email)
	if user != nil {
		utils.ResponseBadRequest(ctx, "Email already registered")
		return
	}

	err := c.authService.RegisterUser(requestBody)
	if err != nil {
		utils.ResponseInternalServerError(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var requestBody dto.LoginDTO

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		utils.ResponseBadRequest(ctx, err.Error())
		return
	}

	user, err := c.authService.AuthenticateUser(requestBody.Email, requestBody.Password)
	if err != nil {
		utils.ResponseUnauthorized(ctx)
		return
	}

	token, err := utils.GenerateToken(requestBody.Email)
	if err != nil {
		utils.ResponseInternalServerError(ctx, err.Error())
		return
	}

	responseData := map[string]interface{}{
		"user":  user,
		"token": token,
	}

	utils.ResponseOK(ctx, responseData)
}
