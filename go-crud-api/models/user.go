package models

type User struct {
	ID       int    `json:"idusuario"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Senha    string `json:"senha"`
	Telefone string `json:"telefone"`
}
