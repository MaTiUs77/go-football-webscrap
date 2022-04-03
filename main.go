package main

import (
	"go-football-webscrap/src"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/bundesliga", src.Bundesliga)
	r.GET("/eurocopa", src.Eurocopa)
	r.GET("/champions_league", src.ChampionLeague)
	r.GET("/premier_league", src.PremierLeague)
	r.GET("/copa_argentina", src.CopaArgentina)
	r.GET("/primera_division_argentina", src.PrimeraDivisionArgentina)

	r.GET("/qatar", src.Qatar)

	r.Run() // listen and serve on 0.0.0.0:8080
}
