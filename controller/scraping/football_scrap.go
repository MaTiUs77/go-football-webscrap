package scraping

import (
	"github.com/gin-gonic/gin"
)

var (
	eurocopa                   = "https://ar.marca.com/claro/futbol-internacional/eurocopa/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	bundesliga                 = "https://ar.marca.com/claro/futbol-internacional/bundesliga/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	champions_league           = "https://ar.marca.com/claro/futbol-internacional/champions-league/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	premier_league             = "https://ar.marca.com/claro/futbol-internacional/premier-league/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	copa_argentina             = "https://ar.marca.com/claro/futbol/copa-argentina/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	primera_division_argentina = "https://ar.marca.com/claro/futbol/primera-division/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
)

func Eurocopa(c *gin.Context) {
	serv(DoScrap(eurocopa), c)
}

func Bundesliga(c *gin.Context) {
	serv(DoScrap(bundesliga), c)
}

func ChampionLeague(c *gin.Context) {
	serv(DoScrap(champions_league), c)
}

func PremierLeague(c *gin.Context) {
	serv(DoScrap(premier_league), c)
}

func CopaArgentina(c *gin.Context) {
	serv(DoScrap(copa_argentina), c)
}

func PrimeraDivisionArgentina(c *gin.Context) {
	serv(DoScrap(primera_division_argentina), c)
}

func serv(evento []Evento, c *gin.Context) {
	c.JSON(200, gin.H{"result": evento})
}
