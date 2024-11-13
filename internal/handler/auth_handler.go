package handler

import (
	"go-api-gin/internal/domain"
	"go-api-gin/internal/service"
	"go-api-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.UserService
}

// Constructor
func NewAuthHandler(service *service.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Método de controlador para Log-in
func (h *AuthHandler) Login(c *gin.Context) {
	var loginDetails domain.LoginRequest

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login request"})
		return
	}

	// Verificar si el usuario existe y obtener datos
	user, err := h.service.GetUserByEmail(loginDetails.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Verificación de contraseña
	if !utils.CheckPasswordHash(loginDetails.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generar el token usando el userID en string
	token, err := utils.GenerateToken(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
