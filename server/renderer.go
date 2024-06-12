package server

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
)

type templRenderer struct{}

func (t *templRenderer) Instance(name string, data any) render.Render {
	c, ok := data.(templ.Component)
	if !ok {
		c = nil
	}
	return renderer{
		ctx:       context.Background(),
		component: c,
	}
}

type renderer struct {
	ctx       context.Context
	component templ.Component
}

func (r renderer) Render(w http.ResponseWriter) error {
	if r.component == nil {
		return nil
	}
	return r.component.Render(r.ctx, w)
}

func (r renderer) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
