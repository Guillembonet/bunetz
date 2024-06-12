package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillembonet/bunetz/views/about_website"
	"github.com/guillembonet/bunetz/views/blog"
	"github.com/guillembonet/bunetz/views/echo"
)

type Static struct {
}

func NewStatic() *Static {
	return &Static{}
}

func (*Static) Blog(c *gin.Context) {
	c.HTML(http.StatusOK, "", blog.Blog())
}

func (*Static) AboutWebsite(c *gin.Context) {
	c.HTML(http.StatusOK, "", about_website.AboutWebsite())
}

func (*Static) Echo(c *gin.Context) {
	echoValue, ok := c.GetQuery("echo")
	if !ok {
		echoValue = "Use the query parameter 'echo' to see the value echoed back."
	}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered:  []string{"application/json", "text/html"},
		HTMLData: echo.Echo(echoValue),
		JSONData: gin.H{"echo": echoValue},
	})
}

func (s *Static) Register(r *gin.RouterGroup) {
	r.GET("/", s.Blog)
	r.GET("/blog", s.Blog)
	r.GET("/about-this-website", s.AboutWebsite)
	r.GET("/echo", s.Echo)
}
