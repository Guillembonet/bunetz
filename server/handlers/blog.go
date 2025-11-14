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
			server.WithBase(b.notFound, "Not found", "", templ.WithStatus(http.StatusNotFound)).ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/blog", http.StatusTemporaryRedirect)
	})
}

func (b *Blog) Blog() http.Handler {
	return server.WithBase(b.blog, "Bunetz's Blog",
		"Blog posts about various topics related to software engineering.")
}

func (b *Blog) BlogPosts() http.Handler {
	return templ.Handler(blog.BlogPostsCards(blog_posts.GetLiveBlogPosts()))
}

func (b *Blog) BlogPost() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		blogPost, htmlContent, err := blog_posts.GetLiveBlogPost(id)
		if err != nil {
			if errors.Is(err, blog_posts.ErrPostNotFound) {
				server.WithBase(b.notFound, "Not found", "", templ.WithStatus(http.StatusNotFound)).ServeHTTP(w, r)
				return
			}
			slog.Error("failed to get blog post", slog.String("error", err.Error()), slog.String("id", id))
			server.WithBase(b.internalServerError, "Internal server error", "", templ.WithStatus(http.StatusInternalServerError)).ServeHTTP(w, r)
			return
		}
		server.WithBase(blog.Post(blogPost.Title, htmlContent), blogPost.Title, blogPost.Description).ServeHTTP(w, r)
	})
}

func (s *Blog) Register(mux *http.ServeMux) {
	mux.Handle("GET /", s.Home())
	mux.Handle("GET /blog", s.Blog())
	mux.Handle("GET /blog/posts", s.BlogPosts())
	mux.Handle("GET /blog/posts/{id}", s.BlogPost())
}
