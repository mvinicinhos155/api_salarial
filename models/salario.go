package models

type Salario struct {
	ID int `json:"id"`
	Tipo string `json:"tipo"`
	Valor float64 `json:"valor"`
	User_id int `json:"user_id"`
}