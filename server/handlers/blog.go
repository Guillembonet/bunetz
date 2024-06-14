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

var blogTitle = "Blog"

type Blog struct {
}

func NewBlog() *Blog {
	return &Blog{}
}

func (*Blog) Blog(c *gin.Context) {
	c.HTML(http.StatusOK, "", server.WithBase(c, blog.Blog(blog_posts.BlogPosts), &blogTitle))
}

func (*Blog) BlogPost(c *gin.Context) {
	id := c.Param("id")
	blogPost, ok := blog_posts.BlogPostsByID[id]
	if !ok {
		c.HTML(http.StatusNotFound, "", server.WithBase(c, error_pages.NotFound(), nil))
		return
	}
	htmlContent, err := blog_posts.GetBlogPostHtml(id)
	if err != nil {
		if errors.Is(err, blog_posts.ErrPostNotFound) {
			c.HTML(http.StatusNotFound, "", server.WithBase(c, error_pages.NotFound(), nil))
			return
		}
		//TODO: return an error page
		c.HTML(http.StatusNotFound, "", server.WithBase(c, error_pages.InternalServerError(), nil))
		return
	}
	c.HTML(http.StatusOK, "", server.WithBase(c, blog.Post(blogPost.Title, htmlContent), &blogPost.Title))
}

func (s *Blog) Register(r *gin.RouterGroup) {
	r.GET("/", s.Blog)
	r.GET("/blog", s.Blog)
	r.GET("/blog/posts/:id", s.BlogPost)
}
