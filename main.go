package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1. MODEL / STRUCT PRODUCT
type Product struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(255)" json:"name"`
	Price int64  `json:"price"`
	Stock int    `json:"stock"`
}

// Global variable untuk Database
var DB *gorm.DB

// 2. FUNGSIONALITAS KONEKSI DATABASE VIA .ENV
func initDB() {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, membaca dari System Environment")
	}

	// Ambil data credential dari file .env
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	// Menyusun DSN (Data Source Name) secara dinamis
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		dbUser, dbPass, dbHost, dbPort, dbName)

	// Membuka koneksi ke MySQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terkoneksi ke database: ", err)
	}

	fmt.Println("Database berhasil terkoneksi dengan aman via .env!")

	// Otomatis sinkronisasi tabel (Auto Migration)
	DB.AutoMigrate(&Product{})
}

// 3. MAIN FUNCTION & ROUTING CONTROLLER
func main() {
	// Set Gin ke mode Release biar performa maksimal di VPS
	gin.SetMode(gin.ReleaseMode)

	// Jalankan koneksi DB
	initDB()

	// Inisialisasi router Gin
	r := gin.Default()

	// ----------------------------------------------------
	// ENDPOINT: GET ALL PRODUCTS (Read)
	// ----------------------------------------------------
	r.GET("/products", func(c *gin.Context) {
		var products []Product
		DB.Find(&products)
		c.JSON(http.StatusOK, gin.H{"data": products})
	})

	// ----------------------------------------------------
	// ENDPOINT: GET SINGLE PRODUCT BY ID (Read Detail)
	// ----------------------------------------------------
	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product Product
		if err := DB.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": product})
	})

	// ----------------------------------------------------
	// ENDPOINT: POST CREATE PRODUCT (Create)
	// ----------------------------------------------------
	r.POST("/products", func(c *gin.Context) {
		var input Product
		// Validasi inputan payload body JSON
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		DB.Create(&input)
		c.JSON(http.StatusCreated, gin.H{"message": "Produk berhasil ditambahkan", "data": input})
	})

	// ----------------------------------------------------
	// ENDPOINT: PUT UPDATE PRODUCT BY ID (Update)
	// ----------------------------------------------------
	r.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product Product
		// Pastikan datanya ada dulu di DB
		if err := DB.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}

		var input Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Melakukan update hanya pada field yang dikirim di JSON
		DB.Model(&product).Updates(input)
		c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil diupdate", "data": product})
	})

	// ----------------------------------------------------
	// ENDPOINT: DELETE PRODUCT BY ID (Delete)
	// ----------------------------------------------------
	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product Product
		// Pastikan datanya ada dulu di DB
		if err := DB.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}

		DB.Delete(&product)
		c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
	})

	// Berjalan secara internal di port 8080 VPS
	r.Run(":8080")
}