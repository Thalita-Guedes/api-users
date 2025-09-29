package models

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	RG      string `json:"rg"`
	CPF     string `json:"cpf"`
	Address string `json:"address"`
}
