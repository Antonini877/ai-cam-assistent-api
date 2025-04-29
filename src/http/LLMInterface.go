package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// LLMConfig defines the configuration for the LLM service
type LLMConfig struct {
	Endpoint   string
	APIKey     string
	Headers    map[string]string
	HTTPClient *http.Client
}

// LLMWrapper is a wrapper for interacting with an LLM service
type LLMWrapper struct {
	config LLMConfig
}

// NewLLMWrapper creates a new instance of LLMWrapper
func NewLLMWrapper(config LLMConfig) *LLMWrapper {
	if config.HTTPClient == nil {
		config.HTTPClient = http.DefaultClient
	}
	return &LLMWrapper{config: config}
}

// ProcessImage sends an image to the LLM service and retrieves the result
func (w *LLMWrapper) ProcessImage(imageData []byte, additionalParams map[string]interface{}) (string, error) {
	// Prepare the request payload
	payload := map[string]interface{}{
		"image": imageData,
	}
	for key, value := range additionalParams {
		payload[key] = value
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", w.config.Endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	if w.config.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+w.config.APIKey)
	}
	for key, value := range w.config.Headers {
		req.Header.Set(key, value)
	}

	// Send the request
	resp, err := w.config.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New("failed to process image: " + string(bodyBytes))
	}

	// Parse the response
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return "", err
	}

	// Extract the result
	result, ok := response["result"].(string)
	if !ok {
		return "", errors.New("invalid response format")
	}

	return result, nil
}