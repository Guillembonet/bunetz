package server

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guillembonet/bunetz/blog_posts"
	"github.com/guillembonet/bunetz/server/middleware"
	"github.com/guillembonet/bunetz/views/assets"
	"github.com/guillembonet/bunetz/views/error_pages"
)

type Server struct {
	server *http.Server
}

type Handler interface {
	Register(*gin.RouterGroup)
}

func NewServer(addr string, handler ...Handler) (*Server, error) {
	g := gin.New()
	g.Use(middleware.Logger, gin.Recovery(), middleware.AssetsCache)
	g.HTMLRender = &templRenderer{}

	g.StaticFS("/assets", http.FS(assets.Assets))

	blogPostsAssets, err := fs.Sub(blog_posts.BlogPostAssets, "assets")
	if err != nil {
		return nil, fmt.Errorf("failed to get blog posts assets: %w", err)
	}
	g.StaticFS("/blog/assets", http.FS(blogPostsAssets))

	rg := g.Group("/")

	g.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "", WithBase(c, error_pages.NotFound(), "Not found", ""))
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
