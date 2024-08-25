package controllers

import (
    "net/http"
    "os"
    "io"
    "src/models"
    "src/views"
    "github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
    file, err := c.FormFile("image")
    if err != nil {
        views.RenderJSON(c, http.StatusBadRequest, gin.H{"error": "Failed to retrieve the image"})
        return
    }

    // Salvar a imagem no servidor (opcional, para processamento local)
    filePath := "temp_image.png"
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        views.RenderJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to save the image"})
        return
    }

    // Chamar a função para enviar a imagem para a API do GPT-4
    result, err := models.ProcessImage(filePath)
    if err != nil {
        views.RenderJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to process the image"})
        return
    }

    views.RenderJSON(c, http.StatusOK, gin.H{"result": result})
}
