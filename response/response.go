package response

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"path/filepath"
)

type HTMLRenderer struct {
	templates map[string]*template.Template
	logger    *slog.Logger
}

func NewHTMLRenderer(dir string, logger *slog.Logger) (*HTMLRenderer, error) {
	pages, err := filepath.Glob(filepath.Join(dir, "*.html"))
	if err != nil {
		return nil, err
	}
	cache := make(map[string]*template.Template, len(pages))
	for _, page := range pages {
		t, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[filepath.Base(page)] = t
	}
	return &HTMLRenderer{templates: cache, logger: logger}, nil
}

func (r *HTMLRenderer) RenderHTML(w http.ResponseWriter, name string, data any) {
	t, ok := r.templates[name]
	if !ok {
		r.ServerError(w, fmt.Errorf("template %q not found", name))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.Execute(w, data); err != nil {
		r.logger.Error("template execute", "template", name, "error", err)
	}
}

// ServerError logs the real error and sends a generic 500 to the client.
func (rs *HTMLRenderer) ServerError(w http.ResponseWriter, err error) {
	rs.logger.Error("server error", "error", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
