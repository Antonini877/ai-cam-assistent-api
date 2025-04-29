package controllers

import (
	"net/http"
	"os"

	"github.com/Antonini877/ai-cam-assistent-api/services"
	"github.com/Antonini877/ai-cam-assistent-api/views"
	"github.com/gin-gonic/gin"
)

// UploadImage processes the uploaded image and returns its description
// Function that handles the upload of an image, processes it using a description service, and returns the result.
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image") // Retrieves the file sent in the form.
	if err != nil {
		views.RenderJSON(c, http.StatusBadRequest, gin.H{"error": "Failed to retrieve the image"}) // Returns an error if the file is not found.
		return
	}

	// Save the image on the server (optional, for local processing)
	// Saves the uploaded file to a temporary path on the server.
	filePath := "temp_image.png"
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		views.RenderJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to save the image"}) // Returns an error if the file cannot be saved.
		return
	}

	// Instantiate the image description service
	// Creates an instance of the image description service using the Hugging Face API.
	apiURL := "https://api-inference.huggingface.co/models/Salesforce/blip-image-captioning"

	imageService := services.ImageDescriptionServiceImpl(apiURL)

	// Process the image and get the description
	// Sends the image to the service and retrieves the generated description.
	result, err := imageService.DescribeImage(filePath)
	if err != nil {
		views.RenderJSON(c, http.StatusInternalServerError, gin.H{"error": "Failed to process the image"}) // Returns an error if processing fails.
		return
	}

	// Return the result
	// Returns the generated description as a JSON response.
	views.RenderJSON(c, http.StatusOK, gin.H{"description": result})
}
