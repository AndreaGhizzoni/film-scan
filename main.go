package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AndreaGhizzoni/film-scan/model"
)

func parseError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	path := "/home/andrea/infos/2012.mkv.json"
	raw, err := ioutil.ReadFile(path)
	parseError(err)

	var f model.FilmStat
	err = json.Unmarshal(raw, &f)
	parseError(err)
}
