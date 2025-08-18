package middleware

import (
  "api/utils"
  "net/http"
  "strings"

  "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    header := c.GetHeader("Authorization")
    if header == "" {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication header"})
      c.Abort()
      return
    }

    parts := strings.Split(header, " ")
    if len(parts) != 2 || parts[0] != "Bearer" {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth header format"})
      c.Abort()
      return
    }

    token, err := utils.ValidateToken(parts[1])
    if err != nil || !token.Valid {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
      c.Abort()
      return
    }

    c.Next()
  }
}
