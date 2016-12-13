// main.go

package main

import (
	"io/ioutil"

	"github.com/AndreaGhizzoni/film-scan/model"
	"github.com/AndreaGhizzoni/film-scan/templates"
	"github.com/gin-gonic/gin"
)

var (
	route *gin.Engine
	films []model.Film
)

// this function parse all json data of movies
func parseAllFiles() {
	basePath := "/home/andrea/infos/" // TODO change this
	files, _ := ioutil.ReadDir(basePath)
	for _, f := range files {
		err, f := model.FromJSON(basePath + f.Name())
		if err != nil {
			panic(err)
		}

		films = append(films, f)
	}
}

func main() {
	parseAllFiles()

	// Set the router as the default one provided by Gin
	route = gin.Default()

	//route.Static("static/css", "/static/css")
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	route.LoadHTMLGlob(templates.BasePath)

	initRoute(route) // in route.go

	// Start serving the application
	route.Run()
}
