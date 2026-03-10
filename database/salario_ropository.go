package database

import (
	"api_salarial/models"
	"database/sql"
	"log"
)

 func InsertSalario (salario *models.Salario, db *sql.DB) error {
	query := `INSERT INTO salario (tipo , valor, user_id) VALUES ($1 , $2, $3)`

	_, err := db.Exec(query, salario.Tipo, salario.Valor, salario.User_id)
		if err != nil {
			log.Printf("Erro ao inserir valor")
			return err
		} else {
			log.Printf("Dados enviado com sucesso")
		}

		return nil
}

func GetTotalSalario (db *sql.DB, user_id int) (float64, error) {
	query := `
	SELECT 
	COALESCE(SUM(CASE WHEN TRIM(tipo) = '+' THEN valor ELSE 0 END),0) -
	COALESCE(SUM(CASE WHEN TRIM(tipo) = '-' THEN valor ELSE 0 END),0)
	AS saldo
	FROM salario
	WHERE user_id = $1
	`

	var saldo float64

	err := db.QueryRow(query, user_id).Scan(&saldo)
		if err != nil {
			return 0, err
		}


	return saldo, nil

}

func GetOneSalario (db *sql.DB, user_id int) ([]models.Salario, error) {

	var salarios []models.Salario 

	query := "SELECT id, tipo, valor, user_id FROM salario WHERE user_id = $1"

	rows ,err := db.Query(query, user_id)
		if err != nil {
			return salarios, err
		}

	defer rows.Close()

	for rows.Next() {
		var salario models.Salario

		err := rows.Scan(&salario.ID, &salario.Tipo, &salario.Valor, &salario.User_id)
			if err != nil {
				return salarios, err
			}

		salarios = append(salarios, salario)
	}

	return salarios, nil
}


