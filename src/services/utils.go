package services

import (
    "bytes"
    "errors"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "encoding/json"
)

type GPT4Response struct {
    Description string `json:"description"`
}

// ProcessImage envia a imagem para a API de IA e retorna a descrição
func ProcessImage(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    var b bytes.Buffer
    writer := multipart.NewWriter(&b)
    part, err := writer.CreateFormFile("image", filePath)
    if err != nil {
        return "", err
    }
    _, err = io.Copy(part, file)
    if err != nil {
        return "", err
    }
    writer.Close()

    // Enviar a imagem para a API de IA (substitua pela URL da API real)
    req, err := http.NewRequest("POST", "https://api.gpt-4.com/process-image", &b)
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", errors.New("failed to get valid response from AI API")
    }

    var gpt4Response GPT4Response
    if err := json.NewDecoder(resp.Body).Decode(&gpt4Response); err != nil {
        return "", err
    }

    return gpt4Response.Description, nil
}
