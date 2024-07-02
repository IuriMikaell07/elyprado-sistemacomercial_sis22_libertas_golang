package models

type Marca struct {
	IDmarca      int    `json:"idmarca"`
	Nomemarca    string `json:"nomemarca"`
	Logo         string `json:"logo"`
	Pais_origem  string `json:"pais_origem"`
	Telefone_sac string `json:"telefone_sac"`
}
