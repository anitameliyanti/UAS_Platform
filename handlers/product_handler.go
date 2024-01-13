package handlers

import (
	"fmt"
	"net/http"
	"uas_platform/database"
	"uas_platform/models"

	"github.com/gin-gonic/gin"
)

// CreateProductHandler Pembuatan produk baru
// @Summary Membuat produk baru
// @Description Membuat produk baru
// @Accept json
// @Produce json
// @Param product body models.Product true "Produk yang akan dibuat"
// @Success 201 {object} models.Product
// @Router /api/products [post]
func CreateProductHandler(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&product).Error; err != nil {
		fmt.Println("Error creating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memasukkan produk"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Produk berhasil dibuat", "product": product})
}

// GetAllProductsHandler Menampilkan semua produk
// @Summary Tampil semua produk
// @Description Daftar semua produk
// @Produce json
// @Success 200 {array} models.Product
// @Router /api/products [get]
func GetAllProductsHandler(c *gin.Context) {
	var products []models.Product

	if err := database.DB.Find(&products).Error; err != nil {
		fmt.Println("Error fetching products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kesalahan Server"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProductHandler Update produk berdasarkan ID
// @Summary Update Produk
// @Description Merubah atau update produk
// @Accept json
// @Produce json
// @Param id path int true "ID Produk"
// @Param product body models.Product true "Update Produk"
// @Success 200 {object} models.Product
// @Router /api/products/{id} [put]
func UpdateProductHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		fmt.Println("Error finding product:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&product).Error; err != nil {
		fmt.Println("Error updating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update produk tidak berhasil"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProductHandler Menghapus produk menggunakan ID
// @Summary Hapus produk
// @Description Hapus produk menggunakan ID
// @Produce json
// @Param id path int true "ID Produk"
// @Success 200 {string} string
// @Router /api/products/{id} [delete]
func DeleteProductHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		fmt.Println("Error finding product:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		fmt.Println("Error deleting product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}
