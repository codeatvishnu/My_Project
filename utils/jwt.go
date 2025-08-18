package utils

import (
    "time"
	"fmt"
    "github.com/golang-jwt/jwt/v4"
)

var JwtSecret = []byte("your-strong-secret") // use env variables in production

func CreateToken(username string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JwtSecret)
}
func ValidateToken(tokenString string) (*jwt.Token, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if token.Method != jwt.SigningMethodHS256 {
      return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }
    return JwtSecret, nil
  })
  return token, err
}
