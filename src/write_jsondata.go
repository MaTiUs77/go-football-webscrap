package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func WriteJsonData(filename string, evento []Jornadas) {
	filename = strings.ToLower(filename) + ".json"
	fmt.Println("[+] Saving result to file:", filename)

	file, _ := json.MarshalIndent(evento, "", " ")
	_ = ioutil.WriteFile("jsondata/"+filename, file, 0644)
}
