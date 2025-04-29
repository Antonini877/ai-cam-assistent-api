package services

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/Antonini877/ai-cam-assistent-api/http" // Import the LLMInterface package
)

// ImageDescriptionService defines an interface for describing images
// Interface that must be implemented by any service that describes images.
type ImageDescriptionService interface {
	DescribeImage(filePath string) (string, error) // Method to describe an image from a file path.
}

// ImageDescriptionServiceImpl is a concrete implementation of ImageDescriptionService
type ImageDescriptionServiceImpl struct {
	LLMWrapper *http.LLMWrapper
}

// DescribeImage sends the image to the LLMWrapper and returns the description
// Method that uploads an image to the Hugging Face API and returns the generated description.
func (s *ImageDescriptionServiceImpl) DescribeImage(filePath string) (string, error) {
	file, err := os.Open(filePath) // Opens the image file.
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the file content and encode it as base64
	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %w", err)
	}
	fileSize := fileInfo.Size()
	if fileSize <= 0 {
		return "", fmt.Errorf("file size is invalid: %d bytes", fileSize)
	}
	fileBuffer := make([]byte, fileSize)
	_, err = file.Read(fileBuffer)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	encodedImage := base64.StdEncoding.EncodeToString(fileBuffer)

	// Use LLMWrapper to process the image
	additionalParams := map[string]interface{}{}
	description, err := s.LLMWrapper.ProcessImage([]byte(encodedImage), additionalParams)
	if err != nil {
		return "", fmt.Errorf("failed to process image: %w", err)
	}

	return description, nil
}
