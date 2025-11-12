package server

import (
	"compress/gzip"
	"context"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/guillembonet/bunetz/blog_posts"
	"github.com/guillembonet/bunetz/server/middleware"
	"github.com/guillembonet/bunetz/views"
)

type Server struct {
	server *http.Server
}

type Handler interface {
	Register(*http.ServeMux)
}

func NewServer(addr string, handler ...Handler) (*Server, error) {
	mux := http.NewServeMux()

	assetsFs := NewNeuteredFileSystem(http.FS(views.Assets))
	mux.Handle("GET /assets/", middleware.AssetsCache(http.FileServer(assetsFs)))

	blogPostsAssetsFs := NewNeuteredFileSystem(http.FS(blog_posts.BlogPostAssets))
	mux.Handle("GET /blog/assets/", http.StripPrefix("/blog", middleware.AssetsCache(http.FileServer(blogPostsAssetsFs))))

	http.StripPrefix("/blog", mux)
	for _, h := range handler {
		h.Register(mux)
	}

	compressor := chimiddleware.NewCompressor(gzip.DefaultCompression)

	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: chimiddleware.Recoverer(middleware.Logger(compressor.Handler(mux))),
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
