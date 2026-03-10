package handlers

import (
	"api_salarial/database"
	"api_salarial/models"
	"api_salarial/token"
	"database/sql"

	"net/http"
	"net/mail"
	"regexp"
    "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HandlerCreateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro ao enviar os dados",
			})
			return
		}

		// Validar email
		if _, err := mail.ParseAddress(user.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Email inválido, tente novamente",
			})
			return
		}

		// Validar senha
		hasUpper := regexp.MustCompile(`[A-Z]`)
		hasNumber := regexp.MustCompile(`[0-9]`)

		if len(user.Password) < 8 ||
			!hasUpper.MatchString(user.Password) ||
			!hasNumber.MatchString(user.Password) {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Senha deve ter no mínimo 8 caracteres, uma letra maiúscula e um número",
			})
			return
		}

		// Hash da senha
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Erro interno no servidor",
			})
			return
		}

		user.Password = string(hash)

		if err := database.InsertUser(&user, db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Erro com banco de dados",
			})
			return
		}

		user.Password = ""

		c.JSON(http.StatusCreated, gin.H{
			"message": "Usuário criado com sucesso",
			"user":    user,
		})
	}
}

func HandlerGetAllUser(db *sql.DB)gin.HandlerFunc{
	return func(c *gin.Context) {

		user, err := database.GetAllUsers(db)
		   if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error" : "Erro ao buscar users"})
			return 
		   }

		   c.JSON(201, gin.H{
			"message" : "Usuário listado com sucesso",
		    "users" : user,
	})
	}
}


func HandlerLogin(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var loginUser struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Bind JSON
		if err := c.ShouldBindJSON(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro ao enviar dados",
			})
			return
		}

		// Buscar usuário no banco
		dbUser, err := database.GetUserByEmail(db, loginUser.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Email ou senha inválidos",
			})
			return
		}

		// Comparar senha
		if err := bcrypt.CompareHashAndPassword(
			[]byte(dbUser.Password),
			[]byte(loginUser.Password),
		); err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Email ou senha inválidos",
			})
			return
		}

		// Gerar token
		tokenString, err := token.GenerateJwt(dbUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Erro interno ao gerar token",
			})
			return
		}

		// Remover senha da resposta
		dbUser.Password = ""

		c.JSON(http.StatusOK, gin.H{
			"message": "Usuário logado com sucesso",
			"user":    dbUser,
			"token":   tokenString,
		})
	}
}

