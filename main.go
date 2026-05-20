package main

import (
	"project-golang/config"
	"project-golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// 1. Jalankan koneksi DB
	config.InitDB()

	// 2. Setup Routing
	r := routes.SetupRouter()

	// 3. Jalankan server
	r.Run(":8080")
}