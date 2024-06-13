package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guillembonet/bunetz/server/middleware"
	"github.com/guillembonet/bunetz/views/assets"
	"github.com/guillembonet/bunetz/views/not_found"
)

type Server struct {
	server *http.Server
}

type Handler interface {
	Register(*gin.RouterGroup)
}

func NewServer(addr string, handler ...Handler) (*Server, error) {
	g := gin.New()
	g.Use(middleware.Logger, gin.Recovery())
	g.HTMLRender = &templRenderer{}

	g.StaticFS("assets", http.FS(assets.Assets))

	rg := g.Group("/")

	g.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "", WithBase(c, not_found.NotFound()))
	})

	for _, h := range handler {
		h.Register(rg)
	}

	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: g,
		},
	}, nil
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}
