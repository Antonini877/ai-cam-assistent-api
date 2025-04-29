package main

import (
	"github.com/Antonini877/ai-cam-assistent-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Route to check the server status
	// Endpoint that returns the status "up" to indicate that the server is running.
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})

	// Route for image upload
	// Endpoint that allows uploading an image and returns its description.
	r.POST("ai-assistant/v1/upload", controllers.UploadImage) // Calls the UploadImage function from the controller.

	r.Run() // Starts the server on the default port (0.0.0.0:8080 or localhost:8080 on Windows).
	// The default port is 8080, but it can be changed by passing an argument to r.Run().
}
