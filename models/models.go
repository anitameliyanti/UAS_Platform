package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Nama     string `json:"nama"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Product struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	NamaProduk string  `json:"nama_produk"`
	Harga      float64 `json:"harga"`
}
