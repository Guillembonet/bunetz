package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillembonet/bunetz/views/about_me"
)

type Contact struct {
}

func NewContact() *Contact {
	return &Contact{}
}

func (co *Contact) Contact(c *gin.Context) {
	message, ok := c.GetPostForm("message")
	if !ok || message == "" {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "", about_me.Contact(message))
}

func (co *Contact) Register(r *gin.RouterGroup) {
	r.POST("/contact", co.Contact)

}
