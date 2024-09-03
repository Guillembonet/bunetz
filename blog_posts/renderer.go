package blog_posts

import (
	"fmt"
	"io"
	"strings"

	"github.com/russross/blackfriday/v2"
)

const (
	H1Tag = "h1"
	H2Tag = "h2"
	H3Tag = "h3"
	H4Tag = "h4"
	PTag  = "p"
)

var customRenderer = CustomRenderer{
	Renderer: blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags,
	}),
	cssClasses: map[string][]string{
		H1Tag: {"text-4xl", "my-6"},
		H2Tag: {"text-3xl", "my-4"},
		H3Tag: {"text-2xl", "my-3"},
		H4Tag: {"text-xl", "my-2"},
		PTag:  {"my-2"},
	},
}

type CustomRendererCSSClasses struct {
	H1 []string
	H2 []string
	H3 []string
	H4 []string
	P  []string
}

type CustomRenderer struct {
	blackfriday.Renderer
	cssClasses map[string][]string
}

func (r *CustomRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	switch node.Type {
	case blackfriday.Heading:
		switch node.Level {
		case 1:
			r.renderElement(w, H1Tag, r.cssClasses[H1Tag], entering)
		case 2:
			r.renderElement(w, H2Tag, r.cssClasses[H2Tag], entering)
		case 3:
			r.renderElement(w, H3Tag, r.cssClasses[H3Tag], entering)
		case 4:
			r.renderElement(w, H4Tag, r.cssClasses[H4Tag], entering)
		}
		return blackfriday.GoToNext
	case blackfriday.Paragraph:
		r.renderElement(w, PTag, r.cssClasses[PTag], entering)
		return blackfriday.GoToNext
	}

	// Use the default renderer for all other elements
	return r.Renderer.RenderNode(w, node, entering)
}

func (r *CustomRenderer) renderElement(w io.Writer, element string, classes []string, entering bool) {
	if entering {
		w.Write([]byte(r.getEnteringElement(element, classes)))
	} else {
		w.Write([]byte(fmt.Sprintf("</%s>", element)))
	}
}

func (r *CustomRenderer) getEnteringElement(element string, classes []string) string {
	return fmt.Sprintf(`<%s%s>`, element, r.getClassSuffix(classes))
}

func (r *CustomRenderer) getClassSuffix(classes []string) string {
	if len(classes) == 0 {
		return ""
	}

	return fmt.Sprintf(` class="%s"`, strings.Join(classes, " "))
}
