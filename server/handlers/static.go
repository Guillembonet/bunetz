package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/guillembonet/bunetz/server"
	"github.com/guillembonet/bunetz/views/about_me"
	"github.com/guillembonet/bunetz/views/about_website"
	"github.com/guillembonet/bunetz/views/echo"
)

type Static struct {
	aboutMe      templ.Component
	aboutWebsite templ.Component
}

func NewStatic() *Static {
	return &Static{
		aboutMe:      about_me.AboutMe(),
		aboutWebsite: about_website.AboutWebsite(),
	}
}

func (s *Static) AboutWebsite() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.WithBase(r, s.aboutWebsite, "About this website",
			"General information about this website.").Render(r.Context(), w)
	})
}

func (s *Static) AboutMe() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.WithBase(r, s.aboutMe, "About me",
			"Information about me, Guillem Bonet.").Render(r.Context(), w)
	})
}

func (*Static) Echo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		echoValue := r.URL.Query().Get("echo")
		if echoValue == "" {
			echoValue = "Use the query parameter 'echo' to see the value echoed back."
		}
		server.WithBase(r, echo.Echo(echoValue), "Echo", "Echo the value back.").Render(r.Context(), w)
	})
}

func (s *Static) Register(mux *http.ServeMux) {
	mux.Handle("GET /about-this-website", s.AboutWebsite())
	mux.Handle("GET /about-me", s.AboutMe())
	mux.Handle("GET /echo", s.Echo())
}
