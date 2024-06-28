package blog_posts

import (
	"embed"
	"fmt"
	"io"
	"time"

	"github.com/russross/blackfriday/v2"
)

var ErrPostNotFound = fmt.Errorf("post not found")

//go:embed *.md
var postsFiles embed.FS

//go:embed assets/*
var BlogPostAssets embed.FS

type BlogPost struct {
	ID               string
	Title            string
	Description      string
	Date             time.Time
	PreviewImagePath string
	PreviewImageAlt  string
	Live             bool
}

var BlogPosts = []BlogPost{
	{
		ID:               "how-i-over-engineered-my-cluster-part-1",
		Title:            "How I over-engineered my Home Kubernetes Cluster: part 1",
		Description:      "Part 1 of a series of posts about how I overengineered my home Kubernetes cluster.",
		PreviewImagePath: "/blog/assets/kubernetes.png",
		PreviewImageAlt:  "Kubernetes logo",
		Date:             time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC),
		Live:             false,
	},
}

var BlogPostsByID = make(map[string]BlogPost, len(BlogPosts))

func GetLiveBlogPost(id string) (*BlogPost, []byte, error) {
	post, ok := BlogPostsByID[id]
	if !ok || !post.Live {
		return nil, nil, ErrPostNotFound
	}
	file, err := postsFiles.Open(fmt.Sprintf("%s.md", id))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read file: %w", err)
	}

	return &post, blackfriday.Run(fileContent, blackfriday.WithRenderer(&customRenderer)), nil
}

func init() {
	for _, post := range BlogPosts {
		BlogPostsByID[post.ID] = post
	}
}
