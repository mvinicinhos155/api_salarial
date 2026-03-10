package handlers

import (
	"api_salarial/database"
	"api_salarial/models"
	"database/sql"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func HandlerCreateSalario(db * sql.DB)gin.HandlerFunc{
	return func(c *gin.Context) {
		userClaims, _ := c.Get("user")
		Claims := userClaims.(*models.Claims)

		userId, err := strconv.Atoi(Claims.ID)
		if err != nil {
			c.JSON(400, gin.H{"error" : "ID invalido"})
			return 
		}


		var input struct{
			Tipo string `json:"tipo"`
			Valor float64 `json:"valor"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro" : "dados invalidos",	
			})
		}

		salario := models.Salario{
			Tipo: input.Tipo,
			Valor: input.Valor,
			User_id: userId,
		}

		if err := database.InsertSalario(&salario, db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro" : "Erro interno com banco de dados",
			})
		}

		c.JSON(201, gin.H{
			"message" : "salario adicionado com sucesso",
			"salario" : salario,
		})
	}
}

func HandlerGetSalario(db *sql.DB)gin.HandlerFunc{
	return func(c *gin.Context) {
		SalarioClaims, _ := c.Get("user")
		Claims := SalarioClaims.(*models.Claims)

		userId, err := strconv.Atoi(Claims.ID)
		 if err != nil {
			c.JSON(400, gin.H{"erro" : "Id invalido"})
		 }


		dbSalario, err := database.GetTotalSalario(db, userId)
			if err != nil {
				c.JSON(400, gin.H{ "erro" : "Erro interno com banco de dados"})
				return 
			}

		c.JSON(200, gin.H{
			"message" : "listando seu salarios",
			"salario" : dbSalario,
		})
	}
} 


func HandlerGetOneSalario (db *sql.DB)gin.HandlerFunc{
	return func(c *gin.Context) {

		SalarioClaims, _ := c.Get("user")
		Claims := SalarioClaims.(*models.Claims)

		userId, err := strconv.Atoi(Claims.ID)
		 if err != nil {
			c.JSON(400, gin.H{"erro" : "Id invalido"})
		 }


		dbSalario, err := database.GetOneSalario(db, userId)
			if err != nil {
				c.JSON(400, gin.H{ "erro" : "Erro interno com banco de dados"})
				return 
			}

		c.JSON(200, gin.H{
			"message" : "listando seu salarios",
			"salario" : dbSalario,
		})
	}
} 

func HandlerGetValores(db *sql.DB)gin.HandlerFunc{
	return func(c *gin.Context) {
		valorClaims, _ := c.Get("user")
		Claims := valorClaims.(*models.Claims)

		userId, err := strconv.Atoi(Claims.ID)
			if err != nil {
				c.JSON(401, gin.H{"erro" : "ID invalido"})
				return 
			}

		total, err := database.GetTotalValores(db, userId)
			if err != nil {
				c.JSON(401, gin.H{"erro" : "Erro interno com o banco de dados"})
			}

		c.JSON(200, gin.H{
			"message" : "Total listado com sucesso",
			"Total" : total,
		})
	}
}