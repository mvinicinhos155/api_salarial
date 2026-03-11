package main

import (
	"api_salarial/database"
	"api_salarial/handlers"
	"api_salarial/servers/middleware"
	"log"
	"os"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {


	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de sistama")
	} 

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	
	r.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message" : "Api funcionando"})
	})

	r.POST("/user", handlers.HandlerCreateUser(db))
	r.GET("/users", middleware.Auth(), middleware.AdminOnly() ,handlers.HandlerGetAllUser(db))
	r.POST("/login", handlers.HandlerLogin(db))

	r.POST("/salario", middleware.Auth(), handlers.HandlerCreateSalario(db))
	r.GET("/salarios", middleware.Auth() ,handlers.HandlerGetSalario(db))
	r.GET("/salario", middleware.Auth(), handlers.HandlerGetOneSalario(db))
	r.GET("/total", middleware.Auth(), handlers.HandlerGetValores(db))

	
	// Roda migrations
	database.Migrations(db)

	// Sobe servidor
	r.Run(":" + port)
}