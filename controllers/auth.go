package controllers

import (
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "api/utils"
)

type User struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func ProtectedHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "This route is protected and accessible only with valid JWT."})
}

func SignUp(c *gin.Context, db *sql.DB) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    query := `INSERT INTO users (username, password) VALUES ($1, $2)`
    if _, err = db.Exec(query, user.Username, string(hashed)); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username may already exist"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": user.Username + " signed up successfully"})
}

func SignIn(c *gin.Context, db *sql.DB) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var storedHash string
    err := db.QueryRow("SELECT password FROM users WHERE username = $1", user.Username).Scan(&storedHash)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(user.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := utils.CreateToken(user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
