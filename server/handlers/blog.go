package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillembonet/bunetz/blog_posts"
	"github.com/guillembonet/bunetz/server"
	"github.com/guillembonet/bunetz/views/blog"
	"github.com/guillembonet/bunetz/views/error_pages"
)

type Blog struct {
}

func NewBlog() *Blog {
	return &Blog{}
}

func (*Blog) Home(c *gin.Context) {
	c.Redirect(http.StatusFound, "/blog")
}

func (*Blog) Blog(c *gin.Context) {
	c.HTML(http.StatusOK, "", server.WithBase(c, blog.Blog(blog_posts.BlogPosts), "Bunetz blog",
		"Blog posts about various topics related to software engineering."))
}

func (*Blog) BlogPost(c *gin.Context) {
	id := c.Param("id")
	blogPost, ok := blog_posts.BlogPostsByID[id]
	if !ok {
		c.HTML(http.StatusNotFound, "", server.WithBase(c, error_pages.NotFound(), "Not found", ""))
		return
	}
	htmlContent, err := blog_posts.GetLiveBlogPostHtml(id)
	if err != nil {
		if errors.Is(err, blog_posts.ErrPostNotFound) {
			c.HTML(http.StatusNotFound, "", server.WithBase(c, error_pages.NotFound(), "Not found", ""))
			return
		}
		c.HTML(http.StatusInternalServerError, "", server.WithBase(c, error_pages.InternalServerError(), "Internal server error", ""))
		return
	}
	c.HTML(http.StatusOK, "", server.WithBase(c, blog.Post(blogPost.Title, htmlContent), blogPost.Title, blogPost.Description))
}

func (s *Blog) Register(r *gin.RouterGroup) {
	r.GET("/", s.Home)
	r.GET("/blog", s.Blog)
	r.GET("/blog/posts/:id", s.BlogPost)
}
