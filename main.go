package main

import (
	"uas_platform/database"
	"uas_platform/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "uas_platform/docs"
)

// @title API UAS Pemrograman Berbasis Platform
// @description API CRUD menggunakan bahasa GO - Anita Meliyanti 2110018
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http https
// @tags Users
// @tags Products
var _ = swaggerFiles.Handler

func main() {
	// Inisialisasi koneksi database
	database.InitDB()

	// Membuat router baru
	r := gin.Default()

	// Middleware CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	// Grouping untuk API
	apiGroup := r.Group("/api")
	{
		// @Tags Users
		// @Summary Membuat user baru
		// @Description Membuat user baru dengan payload JSON
		// @Accept json
		// @Produce json
		// @Param user body User true "User yang akan dibuat"
		// @Success 201 {object} User
		// @Router [post]
		apiGroup.POST("/users", handlers.CreateUserHandler)

		// @Tags Users
		// @Summary Update user
		// @Description update user dengan payload JSON
		// @Accept json
		// @Produce json
		// @Param id path int true "ID Pengguna"
		// @Param user body User true "User yang akan diupdate"
		// @Success 200 {object} User
		// @Router [put]
		apiGroup.PUT("/users/:id", handlers.UpdateUserHandler)

		// @Tags Users
		// @Summary Hapus user
		// @Description Hapus user berdasarkan ID
		// @Produce json
		// @Param id path int true "ID Pengguna"
		// @Success 200 {string} string
		// @Router [delete]
		apiGroup.DELETE("/users/:id", handlers.DeleteUserHandler)

		// @Tags Users
		// @Summary Menampilkan semua pengguna
		// @Description Menampilkan daftar semua pengguna
		// @Produce json
		// @Success 200 {array} User
		// @Router [get]
		apiGroup.GET("/users", handlers.GetAllUsersHandler)

		// @Tags Products
		// @Summary Update produk
		// @Description Perbarui produk dengan payload JSON
		// @Accept json
		// @Produce json
		// @Param id path int true "ID Produk"
		// @Param product body Product true "Objek produk yang akan diupdate"
		// @Success 200 {object} Product
		// @Router [put]
		apiGroup.PUT("/products/:id", handlers.UpdateProductHandler)

		// @Tags Products
		// @Summary Hapus produk berdasarkan ID
		// @Description Hapus produk berdasarkan ID-nya
		// @Produce json
		// @Param id path int true "ID Produk"
		// @Success 200 {string} string
		// @Router [delete]
		apiGroup.DELETE("/products/:id", handlers.DeleteProductHandler)

		// @Tags Products
		// @Summary Membuat produk baru
		// @Description Buat produk baru dengan payload JSON
		// @Accept json
		// @Produce json
		// @Param product body Product true "Objek produk yang akan dibuat"
		// @Success 201 {object} Product
		// @Router [post]
		apiGroup.POST("/products", handlers.CreateProductHandler)

		// @Tags Products
		// @Summary Menampilkan semua produk
		// @Description Menampilkan daftar semua produk
		// @Produce json
		// @Success 200 {array} Product
		// @Router [get]
		apiGroup.GET("/products", handlers.GetAllProductsHandler)

		// @Tags Users
		// @Summary Login user
		// @Description Untuk melakukan login bisa menggunakan get all user untuk melihat semua user
		// @Accept json
		// @Produce json
		// @Param user body User true "Login user"
		// @Success 200 {object} LoginResponse
		// @Router [post]
		apiGroup.POST("/login", handlers.LoginHandler)
	}

	// @Summary Tampilan Dokumentasi API
	// @Description Endpoint UI Swagger untuk API
	// @Produce html
	// @Router /swagger/*any [get]
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	// Menjalankan server
	r.Run(":8080")
}
