package views

import (
    "github.com/gin-gonic/gin"
)

func RenderJSON(c *gin.Context, statusCode int, data interface{}) {
    c.JSON(statusCode, data)
}
