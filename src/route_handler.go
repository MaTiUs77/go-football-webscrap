package src

import (
	"github.com/gin-gonic/gin"
)

var (
	Url_eurocopa                   = "https://ar.marca.com/claro/futbol-internacional/eurocopa/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	Url_bundesliga                 = "https://ar.marca.com/claro/futbol-internacional/bundesliga/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	Url_champions_league           = "https://ar.marca.com/claro/futbol-internacional/champions-league/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	Url_premier_league             = "https://ar.marca.com/claro/futbol-internacional/premier-league/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	Url_copa_argentina             = "https://ar.marca.com/claro/futbol/copa-argentina/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	Url_primera_division_argentina = "https://ar.marca.com/claro/futbol/primera-division/fixture.html?intcmp=MENUMIGA&s_kw=fixture"
	Url_qatar                      = "https://es.wikipedia.org/wiki/Copa_Mundial_de_F%C3%BAtbol_de_2022"
)

func Eurocopa(c *gin.Context) {
	jsonResponse(ScrapMarca(Url_eurocopa), c)
}

func Bundesliga(c *gin.Context) {
	jsonResponse(ScrapMarca(Url_bundesliga), c)
}

func ChampionLeague(c *gin.Context) {
	jsonResponse(ScrapMarca(Url_champions_league), c)
}

func PremierLeague(c *gin.Context) {
	jsonResponse(ScrapMarca(Url_premier_league), c)
}

func CopaArgentina(c *gin.Context) {
	jsonResponse(ScrapMarca(Url_copa_argentina), c)
}

func PrimeraDivisionArgentina(c *gin.Context) {
	jsonResponse(ScrapMarca(Url_primera_division_argentina), c)
}

func Qatar(c *gin.Context) {
	jsonResponse(ScrapWikipedia(Url_qatar), c)
}

func jsonResponse(evento []Jornadas, c *gin.Context) {
	c.JSON(200, gin.H{"result": evento})
}
