package src

type Match struct {
	NroPartido int    `json:"nro_partido"`
	Played     bool   `json:"played"`
	Local      string `json:"local"`
	Resultado  string `json:"resultado"`
	Visitante  string `json:"visitante"`
	Date       string `json:"date"`
	Fase       string `json:"fase"`
	Grupo      string `json:"grupo"`
}

type Jornadas struct {
	Jornada int     `json:"jornada"`
	Matches []Match `json:"matches"`
}
