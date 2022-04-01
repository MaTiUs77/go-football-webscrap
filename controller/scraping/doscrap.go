package scraping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func DoScrap(url string) []Evento {
	evento := []Evento{}
	competencia := "undefined"

	cc := colly.NewCollector()
	cc.OnRequest(func(r *colly.Request) {
		fmt.Println("[+] Visiting:", r.URL)
	})

	cc.OnHTML("table[class=\"jor agendas\"]", func(e *colly.HTMLElement) {
		//table := e.Attr("id")
		newEvento := Evento{}
		matches := []Match{}

		e.ForEach("caption", func(_ int, caption *colly.HTMLElement) {
			regex := regexp.MustCompile(`^[A-Za-z]*\s(\d{1,2})$`)
			res := regex.FindAllStringSubmatch(caption.Text, -1)

			newEvento.Jornada, _ = strconv.Atoi(res[0][1]) // Solo quiero el numero de "Jornada 34"
		})

		e.ForEach("tbody tr", func(tr_idx int, tr *colly.HTMLElement) {
			date := ""
			local := tr.ChildText("td:nth-child(1)")
			resultado := tr.ChildText("td:nth-child(2)")
			visitante := tr.ChildText("td:nth-child(3)")
			played, _ := regexp.MatchString(`^(\d{1,2})-(\d{1,2})$`, resultado)

			if !played {
				date = resultado
				resultado = "0-0"
			}

			match := Match{
				Played:    played,
				Local:     local,
				Resultado: resultado,
				Visitante: visitante,
				Date:      date,
			}

			//fmt.Println("Played", played, "Local", local, "Resultado", resultado, "Visitante", visitante)
			matches = append(matches, match)
			newEvento.Matches = matches
			newEvento.Eventos = len(newEvento.Matches)
		})

		evento = append(evento, newEvento)
	})

	cc.OnHTML("li[class=\"second-level\"] span[itemprop=\"name\"]:nth-child(1)", func(e *colly.HTMLElement) {
		competencia = e.Text
	})

	cc.Visit(url)

	saveResultToFile(competencia, evento)

	fmt.Println("[+] Scraping complete!")
	return evento
}

func saveResultToFile(filename string, evento []Evento) {
	filename = strings.ToLower(filename) + ".json"
	fmt.Println("[+] Saving result to file:", filename)

	file, _ := json.MarshalIndent(evento, "", " ")
	_ = ioutil.WriteFile("jsondata/"+filename, file, 0644)
}
