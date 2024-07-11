package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillembonet/bunetz/external/telegram"
	"github.com/guillembonet/bunetz/views/about_me"
)

type Contact struct {
	telegramClient *telegram.Client
}

func NewContact(telegramClient *telegram.Client) *Contact {
	return &Contact{
		telegramClient: telegramClient,
	}
}

func (co *Contact) Contact(c *gin.Context) {
	name, contact, message := c.PostForm("name"), c.PostForm("contact"), c.PostForm("message")
	if name == "" || contact == "" || message == "" {
		nameError, contactError, messageError := "Name is required", "Contact is required", "Message is required"
		if name != "" {
			nameError = ""
		}
		if contact != "" {
			contactError = ""
		}
		if message != "" {
			messageError = ""
		}
		c.HTML(http.StatusBadRequest, "", about_me.Contact(name, nameError, contact, contactError, message, messageError))
		return
	}

	if err := co.telegramClient.SendMessage(name, contact, message); err != nil {
		c.HTML(http.StatusInternalServerError, "", about_me.Contact(name, "", contact, "", message, "Failed to send message"))
		return
	}

	c.HTML(http.StatusOK, "", about_me.ContactSuccess())
}

func (co *Contact) Register(r *gin.RouterGroup) {
	r.POST("/contact", co.Contact)
}
