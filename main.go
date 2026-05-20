package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1. MODEL (Ibaratnya file Model di Laravel + otomatis jadi Migration)
type Product struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(255)" json:"name"`
	Price int64  `json:"price"`
	Stock int    `json:"stock"`
}

var DB *gorm.DB

// 2. KONEKSI DATABASE
func initDB() {
	// Sesuai dengan konfigurasi DB yang kamu kasih
	dsn := "rizal:rizal@tcp(203.194.113.87:3306)/rizal?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database: ", err)
	}

	fmt.Println("Database berhasil terkoneksi!")

	// Otomatis bikin/update tabel 'products' (Mirip php artisan migrate)
	DB.AutoMigrate(&Product{})
}

func main() {
	// Set Gin ke mode release biar ringan di VPS
	gin.SetMode(gin.ReleaseMode)

	// Inisialisasi DB
	initDB()
	r := gin.Default()
	// 3. ROUTING & CONTROLLER (API ENDPOINTS)
	// [GET] ALL PRODUCTS (Mirip Product::all())
	r.GET("/products", func(c *gin.Context) {
		var products []Product
		DB.Find(&products)
		c.JSON(http.StatusOK, gin.H{"data": products})
	})

	// [GET] SINGLE PRODUCT (Mirip Product::find($id))
	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product Product
		if err := DB.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": product})
	})

	// [POST] CREATE PRODUCT (Mirip Product::create($request->all()))
	r.POST("/products", func(c *gin.Context) {
		var input Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		DB.Create(&input)
		c.JSON(http.StatusCreated, gin.H{"message": "Produk berhasil dibuat", "data": input})
	})

	// [PUT] UPDATE PRODUCT (Mirip $product->update())
	r.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product Product
		if err := DB.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}

		var input Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update field yang dikirim
		DB.Model(&product).Updates(input)
		c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil diupdate", "data": product})
	})

	// [DELETE] DELETE PRODUCT (Mirip $product->delete())
	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product Product
		if err := DB.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}

		DB.Delete(&product)
		c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
	})

	// Jalankan di port 8080 internal VPS
	r.Run(":8080")
}
