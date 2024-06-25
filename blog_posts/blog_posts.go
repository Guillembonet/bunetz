package blog_posts

import (
	"embed"
	"fmt"
	"io"

	"github.com/russross/blackfriday/v2"
)

var ErrPostNotFound = fmt.Errorf("post not found")

//go:embed *.md
var postsFiles embed.FS

//go:embed assets/*
var BlogPostAssets embed.FS

type BlogPost struct {
	ID          string
	Title       string
	Description string
	Live        bool
}

var BlogPosts = []BlogPost{
	{
		ID:          "how-i-overengineered-my-cluster-part-1",
		Title:       "How I overengineered my Home Kubernetes Cluster",
		Description: "A series of posts about how I overengineered my home Kubernetes cluster.",
		Live:        false,
	},
}

var BlogPostsByID = make(map[string]BlogPost, len(BlogPosts))

func GetLiveBlogPostHtml(id string) ([]byte, error) {
	post, ok := BlogPostsByID[id]
	if !ok || !post.Live {
		return nil, ErrPostNotFound
	}
	file, err := postsFiles.Open(fmt.Sprintf("%s.md", id))
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return blackfriday.Run(fileContent, blackfriday.WithRenderer(&customRenderer)), nil
}

func init() {
	for _, post := range BlogPosts {
		BlogPostsByID[post.ID] = post
	}
}
