package routes

import (
	"net/http"

	"auth_service/models"
	"auth_service/services"

	"github.com/gin-gonic/gin"
)

// Используем интерфейс
func AuthRoutes(r *gin.Engine, authService services.AuthServiceInterface) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", func(c *gin.Context) {
			var input models.AuthRequest
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			user := &models.User{
				Email:    input.Email,
				Password: input.Password,
			}

			if err := authService.Register(user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"message": "User created"})
		})

		authGroup.POST("/login", func(c *gin.Context) {
			var input models.AuthRequest
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			token, err := authService.Login(&input)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": token})
		})
	}
}
