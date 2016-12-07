// route.go

package main

import (
	"net/http"

	"github.com/AndreaGhizzoni/film-scan/entry_point"
	"github.com/AndreaGhizzoni/film-scan/model"
	"github.com/AndreaGhizzoni/film-scan/templates"
	"github.com/gin-gonic/gin"
)

func initRoute(route *gin.Engine) {
	route.GET(entry_point.Index, showIndex)
	route.GET(entry_point.ViewMovie, showMovie)
}

// TODO change this doc
// Define the route for the index page and display the index.html template
// To start with, we'll use an inline route handler. Later on, we'll create
// standalone functions that will be used as route handlers.
func showIndex(c *gin.Context) {

	render(
		c,
		gin.H{
			"title":   "Home Page",
			"payload": films},
		templates.Index)
}

// TODO add doc
func showMovie(c *gin.Context) {
	// Check if the movie ID is valid
	movieId := c.Param("movie_id")

	// Check if the movie exists
	if f, err := model.GetMovieByID(movieId, films); err == nil {
		render(
			c,
			gin.H{
				"title":   f.Name,
				"payload": f},
			templates.ViewMovie)
	} else {
		// If the article is not found, abort with an error
		c.AbortWithError(http.StatusNotFound, err)
	}

	/* TODO do this check
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
	*/
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data)
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data)
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
