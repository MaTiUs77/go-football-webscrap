package src

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

/*
	Prototipo, vendria bien una buena refactorizacion.
*/
func ScrapWikipedia(url string) []Jornadas {
	jornadas := []Jornadas{}
	competencia := "Qatar"

	cc := colly.NewCollector()
	cc.OnRequest(func(r *colly.Request) {
		fmt.Println("[+] Visiting:", r.URL)
	})

	iteration := 1
	grupos := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	cc.OnHTML("table.autocollapse tr", func(e *colly.HTMLElement) {
		newJornada := Jornadas{}
		matches := []Match{}

		dom := e.DOM.Children()
		date := strings.TrimSpace(dom.Eq(0).Text())
		local := strings.TrimSpace(dom.Eq(1).Eq(0).Text())
		partido := strings.TrimSpace(dom.Eq(2).Eq(0).Text())
		visitante := strings.TrimSpace(dom.Eq(3).Eq(0).Text())

		if date != "" && visitante != "" {
			regex := regexp.MustCompile(`^[A-Za-z]*\s(\d{1,2})$`)
			res := regex.FindAllStringSubmatch(partido, -1)
			nro_partido, _ := strconv.Atoi(res[0][1]) // Solo quiero 34 de Partido 34

			fase := ""
			grupo := ""

			if nro_partido <= 48 {
				fase = "Grupos"
				grupo = grupos[0]
				if iteration%6 == 0 {
					grupos = grupos[1:]
				}
			}

			if nro_partido >= 49 && nro_partido <= 56 {
				fase = "Octavos"
			}

			if nro_partido >= 57 && nro_partido <= 60 {
				fase = "Cuartos"
			}

			if nro_partido >= 61 && nro_partido <= 62 {
				fase = "Semifinales"
			}

			if nro_partido == 63 {
				fase = "Tercer puesto"
			}

			if nro_partido == 64 {
				fase = "Final"
			}

			date = strings.ReplaceAll(date, " de noviembre,", "/11/2020")
			date = strings.ReplaceAll(date, " de diciembre,", "/12/2020")

			match := Match{
				NroPartido: nro_partido,
				Played:     false,
				Local:      local,
				Resultado:  "0-0",
				Visitante:  visitante,
				Date:       date,
				Fase:       fase,
				Grupo:      grupo,
			}

			matches = append(matches, match)
			newJornada.Matches = matches

			jornadas = append(jornadas, newJornada)
			iteration++
		}
	})

	cc.Visit(url)

	// Orden ascendente segun NroPartido
	sort.Slice(jornadas[:], func(i, j int) bool {
		return jornadas[i].Matches[0].NroPartido < jornadas[j].Matches[0].NroPartido
	})

	jornadas = groupByJourney(jornadas)

	WriteJsonData(competencia, jornadas)

	fmt.Println("[+] Scraping complete!")
	return jornadas
}

func groupByJourney(jornadas []Jornadas) []Jornadas {
	newJornada := []Jornadas{}
	curentDate := ""
	i := 0
	for k, v := range jornadas {
		iterationDate := extractDate(v.Matches[0].Date)
		if iterationDate != curentDate {
			i++
			curentDate = iterationDate

			jornadas[k].Jornada = i
			newJornada = append(newJornada, jornadas[k])
		} else {
			newJornada[i-1].Matches = append(newJornada[i-1].Matches, jornadas[k].Matches[0])
		}
	}

	return newJornada
}

func extractDate(date string) string {
	regex := regexp.MustCompile(`^(\d{1,2}/\d{1,2}/\d{4})\s(\d{1,2}:\d{1,2})$`)
	res := regex.FindAllStringSubmatch(date, -1)
	return res[0][1]
}

func stringToDate(date string) time.Time {
	t, err := time.Parse("02/01/2006 15:04", date)
	if err != nil {
		log.Fatal(err)
	}

	return t
}
