package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillembonet/bunetz/server"
	"github.com/guillembonet/bunetz/views/about_me"
	"github.com/guillembonet/bunetz/views/about_website"
	"github.com/guillembonet/bunetz/views/echo"
)

type Static struct {
}

func NewStatic() *Static {
	return &Static{}
}

func (*Static) AboutWebsite(c *gin.Context) {
	c.HTML(http.StatusOK, "", server.WithBase(c, about_website.AboutWebsite(), "About this website",
		"General information about this website."))
}

func (*Static) AboutMe(c *gin.Context) {
	c.HTML(http.StatusOK, "", server.WithBase(c, about_me.AboutMe(), "About me",
		"Information about me, Guillem Bonet."))
}

func (*Static) Echo(c *gin.Context) {
	echoValue, ok := c.GetQuery("echo")
	if !ok {
		echoValue = "Use the query parameter 'echo' to see the value echoed back."
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered:  []string{"application/json", "text/html"},
		HTMLData: server.WithBase(c, echo.Echo(echoValue), "Echo", "Echo the value back."),
		JSONData: gin.H{"echo": echoValue},
	})
}

func (s *Static) Register(r *gin.RouterGroup) {
	r.GET("/about-this-website", s.AboutWebsite)
	r.GET("/about-me", s.AboutMe)
	r.GET("/echo", s.Echo)
}
