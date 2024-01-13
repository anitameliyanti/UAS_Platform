package handlers

import (
	"net/http"
	"strconv"
	"uas_platform/database"
	"uas_platform/models"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler Pembuatan user baru
// @Summary Create user
// @Description Membuat user baru
// @Accept json
// @Produce json
// @Param user body models.User true "User yang akan dibuat"
// @Success 201 {object} models.User
// @Router /api/users [post]
func CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// UpdateUserHandler Update user dengan ID
// @Summary Update user
// @Description Update user dengan menggunakan ID
// @Accept json
// @Produce json
// @Param id path int true "ID Pengguna"
// @Param user body models.User true "Update user"
// @Success 200 {object} models.User
// @Router /api/users/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	idStr := c.Params.ByName("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uint(id)

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui pengguna"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUserHandler Menghapus user dengan ID
// @Summary Hapus user menggunakan ID
// @Description Hapus user
// @Produce json
// @Param id path int true "ID Pengguna"
// @Success 200 {string} string
// @Router /api/users/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}

// LoginHandler Login user
// @Summary Login user
// @Description Untuk melakukan login bisa menggunakan get all user untuk melihat semua user
// @Accept json
// @Produce json
// @Param loginUser body models.LoginUser true "Login user"
// @Success 200 {object} models.LoginUser
// @Router /api/login [post]
func LoginHandler(c *gin.Context) {
	var loginUser models.LoginUser
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFromDB models.User
	if err := database.DB.Where("username = ? AND password = ?", loginUser.Username, loginUser.Password).First(&userFromDB).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kredensial login tidak valid"})
		return
	}

	// Otentikasi berhasil, tidak perlu membuat token JWT

	c.JSON(http.StatusOK, gin.H{"message": "Login Berhasil"})
}

// GetAllUsersHandler Menampilkan semua user
// @Summary Tampilan semua user
// @Description Menampilkan daftar semua user
// @Produce json
// @Success 200 {array} models.User
// @Router /api/users [get]
func GetAllUsersHandler(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kesalahan Server"})
		return
	}

	c.JSON(http.StatusOK, users)
}
