package database

import (
	"database/sql"
	"log"
)


func CreateTableUser (db *sql.DB) {

	query := `
		CREATE TABLE If NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`


	_, err := db.Exec(query)

		if err != nil {
			log.Printf("Erro ao tentar criar tabela!")
		}

	log.Println("Tabela criada com sucesso!")
}

func InsertRoleTable (db *sql.DB) {
	query := (`
				ALTER TABLE users
				ADD COLUMN IF NOT EXISTS role TEXT NOT NULL DEFAULT 'user'
	`)

	_, err := db.Exec(query)
		if err != nil {
			log.Printf("Erro ao tentar criar role: %v", err)
		} else {
		log.Printf("Role criado com sucesso")	
		}
	
}

func UpdateRole (db *sql.DB) {
	query := `UPDATE users SET role = 'admin' WHERE email = 'vinicius11@gmail.com'`

	_, err := db.Exec(query)
		if err != nil {
			log.Printf("Erro ao tentar atualizar valor no role: %v", err)
		} else {
			log.Printf("Atualização concluido")	
		}
}

func CreateTableSalario(db *sql.DB) {

	query := (`
			CREATE TABLE IF NOT EXISTS salario (
				id SERIAL PRIMARY KEY,
				tipo TEXT NOT NULL,
				valor INT NOT NULL, 
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			)
	`)

	_, err := db.Exec(query)
		if err != nil {
			log.Printf("Erro ao criar tabela")
			return
		} else {
			log.Printf("tabela criada com sucesso")
		}
}

func InsertIdUserTable (db *sql.DB) {
	query := (`
				ALTER TABLE salario
				ADD COLUMN IF NOT EXISTS user_id INTEGER;

				ALTER TABLE salario
				ADD CONSTRAINT salario_user_fk
				FOREIGN KEY (user_id)
				REFERENCES users(id)
				ON DELETE CASCADE;

				ALTER TABLE salario
				ALTER COLUMN user_id SET NOT NULL;
	`)

	_, err := db.Exec(query)
		if err != nil {
			log.Printf("Erro ao tentar criar coluna: %v", err)
		} else {
		log.Printf("Coluna criado com sucesso")	
		}
	
}

func Migrations (db *sql.DB) {
	
}