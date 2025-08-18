package main

import (
    "database/sql"
    "log"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
    "api/controllers"
	"api/middleware"
)

func main() {
    router := gin.Default()

    // PostgreSQL connection string – adjust as needed
    dsn := "host=localhost port=5432 user=postgres password=Fyers@2025 dbname=mydb sslmode=disable"
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Database connection error:", err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }
    log.Println("Database connected!")

    router.POST("/signup", func(c *gin.Context) {
        controllers.SignUp(c, db)
    })
    router.POST("/signin", func(c *gin.Context) {
        controllers.SignIn(c, db)
    })
	protected := router.Group("/")
  protected.Use(middleware.AuthMiddleware())
  {
    protected.GET("/protected", controllers.ProtectedHandler)
  }

    router.Run(":8080")
}
