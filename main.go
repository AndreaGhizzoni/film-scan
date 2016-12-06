package main

import (
	"fmt"
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
	err, f := model.FromJSON(path)
	parseError(err)

	fmt.Println(f.ToJSON())
}
