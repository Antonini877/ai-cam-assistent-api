package models

import (
    "bytes"
    "errors"
    "io"
    "mime/multipart"
    "net/http"
    "os"
)

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
        return "", errors.New("failed to get valid response from GPT-4 API")
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}
