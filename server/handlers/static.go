package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillembonet/bunetz/server"
	"github.com/guillembonet/bunetz/views/about_website"
	"github.com/guillembonet/bunetz/views/echo"
)

type Static struct {
}

func NewStatic() *Static {
	return &Static{}
}

var aboutWebsiteTitle = "About this website"

func (*Static) AboutWebsite(c *gin.Context) {
	c.HTML(http.StatusOK, "", server.WithBase(c, about_website.AboutWebsite(), &aboutWebsiteTitle))
}

var echoTitle = "Echo"

func (*Static) Echo(c *gin.Context) {
	echoValue, ok := c.GetQuery("echo")
	if !ok {
		echoValue = "Use the query parameter 'echo' to see the value echoed back."
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered:  []string{"application/json", "text/html"},
		HTMLData: server.WithBase(c, echo.Echo(echoValue), &echoTitle),
		JSONData: gin.H{"echo": echoValue},
	})
}

func (s *Static) Register(r *gin.RouterGroup) {
	r.GET("/about-this-website", s.AboutWebsite)
	r.GET("/echo", s.Echo)
}
