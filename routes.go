// route.go

package main

import (
	"net/http"

	"github.com/AndreaGhizzoni/film-scan/entry_point"
	"github.com/AndreaGhizzoni/film-scan/model"
	"github.com/AndreaGhizzoni/film-scan/templates"
	"github.com/gin-gonic/gin"
)

// This function initialize all the route
func initRoute(route *gin.Engine) {
	route.GET(entry_point.Index, showIndex)
	route.GET(entry_point.ViewMovie, showMovie)
}

// Define the route for the home page
func showIndex(c *gin.Context) {
	render(
		c,
		gin.H{
			"title":   "Home Page",
			"payload": films,
		},
		templates.Index,
	)
}

// Define the route for viewing a movie stats page
func showMovie(c *gin.Context) {
	// Check if the movie ID is valid
	movieId := c.Param("movie_id")

	// Check if the movie exists
	if f, err := model.GetMovieByID(movieId, films); err == nil {
		render(
			c,
			gin.H{
				"title":   f.Name,
				"payload": f,
			},
			templates.ViewMovie,
		)
	} else {
		// If movieID is not found, abort with an error
		c.AbortWithError(http.StatusNotFound, err)
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that the
// template name is present
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
