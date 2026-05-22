package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller untuk handle halaman utama / index
func IndexPage(c *gin.Context) {
	htmlContent := `
	<!DOCTYPE html>
	<html lang="id">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Central API Golang</title>
		<style>
			body {
				margin: 0;
				padding: 0;
				background: #0d1117;
				color: #58a6ff;
				font-family: 'Courier New', Courier, monospace;
				display: flex;
				justify-content: center;
				align-items: center;
				height: 100vh;
				overflow: hidden;
			}
			.container {
				text-align: center;
				border: 2px solid #238636;
				padding: 40px;
				border-radius: 10px;
				background: #161b22;
				box-shadow: 0 0 20px rgba(35, 134, 54, 0.5);
			}
			h1 {
				color: #238636;
				margin-bottom: 10px;
				font-size: 2.5rem;
			}
			p {
				color: #8b949e;
				font-size: 1.2rem;
			}
			.status {
				display: inline-block;
				padding: 5px 15px;
				background: rgba(35, 134, 54, 0.2);
				color: #3fb950;
				border-radius: 20px;
				font-weight: bold;
				margin-top: 15px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>⚡ Central API Golang ⚡</h1>
			<p>Selamat datang di gate utama sistem backend Golang</p>
			<p style="color: #ff7b72;">Crafted with passion by Rizal</p>
			<div class="status">● SERVICE ONLINE</div>
		</div>
	</body>
	</html>
	`
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
}