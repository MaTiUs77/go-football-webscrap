package main

import (
	"github.com/gin-gonic/gin"

	scraping "rest-gin/controller/scraping"
)

func main() {
	r := gin.Default()
	r.GET("/bundesliga", scraping.Bundesliga)
	r.GET("/eurocopa", scraping.Eurocopa)
	r.GET("/champions_league", scraping.ChampionLeague)
	r.GET("/premier_league", scraping.PremierLeague)
	r.GET("/copa_argentina", scraping.CopaArgentina)
	r.GET("/primera_division_argentina", scraping.PrimeraDivisionArgentina)

	r.Run() // listen and serve on 0.0.0.0:8080
}
