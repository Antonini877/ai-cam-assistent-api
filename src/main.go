package main

import (
    "github.com/gin-gonic/gin"
    "github.com/Antonini877/ai-cam-assistent-api/controllers" 
)

func main() {
    r := gin.Default()

    // Rota para ping
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "up",
        })
    })

    // Rota para upload de imagem
    r.POST("/upload", controllers.UploadImage) // Chama a função UploadImage do controller

    r.Run() // escuta e serve na 0.0.0.0:8080 (para Windows "localhost:8080")
}
