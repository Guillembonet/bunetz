package handlers

import (
	"net/http"

	"github.com/a-h/templ"
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

func (co *Contact) Contact() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, contact, message := r.PostFormValue("name"), r.PostFormValue("contact"), r.PostFormValue("message")
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
			templ.Handler(about_me.Contact(name, nameError, contact, contactError, message, messageError),
				templ.WithStatus(http.StatusBadRequest)).ServeHTTP(w, r)
			return
		}

		if err := co.telegramClient.SendMessage(name, contact, message); err != nil {
			templ.Handler(about_me.Contact(name, "", contact, "", message, "Failed to send message"),
				templ.WithStatus(http.StatusInternalServerError)).ServeHTTP(w, r)
			return
		}

		templ.Handler(about_me.ContactSuccess()).ServeHTTP(w, r)
	})
}

func (co *Contact) Register(mux *http.ServeMux) {
	mux.Handle("POST /contact", co.Contact())
}
