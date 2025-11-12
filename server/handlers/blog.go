package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/guillembonet/bunetz/blog_posts"
	"github.com/guillembonet/bunetz/server"
	"github.com/guillembonet/bunetz/views/blog"
	"github.com/guillembonet/bunetz/views/error_pages"
)

type Blog struct {
	notFound            templ.Component
	internalServerError templ.Component
	blog                templ.Component
}

func NewBlog() *Blog {
	return &Blog{
		notFound:            error_pages.NotFound(),
		internalServerError: error_pages.InternalServerError(),
		blog:                blog.Blog(),
	}
}

func (b *Blog) Home() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			server.WithBase(r, b.notFound, "Not found", "").Render(r.Context(), w)
			return
		}
		http.Redirect(w, r, "/blog", http.StatusTemporaryRedirect)
	})
}

func (b *Blog) Blog() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.WithBase(r, b.blog, "Bunetz's Blog",
			"Blog posts about various topics related to software engineering.").Render(r.Context(), w)
	})
}

func (b *Blog) BlogPosts() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		blog.BlogPostsCards(blog_posts.GetLiveBlogPosts()).Render(r.Context(), w)
	})
}

func (b *Blog) BlogPost() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		blogPost, htmlContent, err := blog_posts.GetLiveBlogPost(id)
		if err != nil {
			if errors.Is(err, blog_posts.ErrPostNotFound) {
				w.WriteHeader(http.StatusNotFound)
				server.WithBase(r, b.notFound, "Not found", "").Render(r.Context(), w)
				return
			}
			slog.Error("failed to get blog post", slog.String("error", err.Error()), slog.String("id", id))
			w.WriteHeader(http.StatusInternalServerError)
			server.WithBase(r, b.internalServerError, "Internal server error", "").Render(r.Context(), w)
			return
		}
		server.WithBase(r, blog.Post(blogPost.Title, htmlContent), blogPost.Title, blogPost.Description).Render(r.Context(), w)
	})
}

func (s *Blog) Register(mux *http.ServeMux) {
	mux.Handle("GET /", s.Home())
	mux.Handle("GET /blog", s.Blog())
	mux.Handle("GET /blog/posts", s.BlogPosts())
	mux.Handle("GET /blog/posts/{id}", s.BlogPost())
}
