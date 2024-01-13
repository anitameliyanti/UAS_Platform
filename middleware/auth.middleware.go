package middleware

import (
	"net/http"

	"uas_platform/database"
	"uas_platform/models"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUser models.LoginUser
		if err := c.ShouldBindJSON(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		var userFromDB models.User
		if err := database.DB.Where("username = ? AND password = ?", loginUser.Username, loginUser.Password).First(&userFromDB).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Kredensial login tidak valid"})
			c.Abort()
			return
		}

		// Set data pengguna ke konteks untuk digunakan di handler
		c.Set("userID", userFromDB.ID)
		c.Next()
	}
}
