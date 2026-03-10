package database

import (
	"api_salarial/models"
	"database/sql"
	"log"
)


func InsertUser (user *models.User, db *sql.DB) error {

	query := "INSERT INTO users (name, email, password) VALUES ($1 , $2 , $3)"

	_, err := db.Exec(query, user.Name, user.Email, user.Password)

		if err != nil {
			return err
		}

		log.Println("Dados Inserido com sucesso!!")
		return nil
}

func GetAllUsers (db *sql.DB) ([]models.User ,error) {

	
	query := ("SELECT id, name, email, password, role FROM users")

	rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
	
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
			if err != nil {
				return nil, err
			}
		
		users = append(users, user)
	}

	return users, nil
}

func GetUserByEmail(db *sql.DB, email string) (models.User, error) {

	var user models.User

	query := "SELECT id, name, email, password, role FROM users WHERE email = $1"

	err := db.QueryRow(query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)

	if err != nil {
		return user, err
	}

	return user, nil
}

func GetTotalValores(db *sql.DB, userId int) (float64, error) {

	var total float64

	query := "SELECT SUM(valor) FROM salario WHERE user_id = $1"

	err := db.QueryRow(query, userId).Scan(&total)
		if err != nil {
			return 0, err
		}

	return total, nil	
}
