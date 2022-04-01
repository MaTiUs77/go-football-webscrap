package scraping

type Match struct {
	Played    bool   `json:"played"`
	Local     string `json:"local"`
	Resultado string `json:"resultado"`
	Visitante string `json:"visitante"`
	Date      string `json:"date"`
}

type Evento struct {
	Jornada int     `json:"jornada"`
	Eventos int     `json:"eventos"`
	Matches []Match `json:"matches"`
}
