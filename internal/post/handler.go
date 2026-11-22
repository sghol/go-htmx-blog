package post

import (
	"net/http"
	"log/slog"
	// "html/template"
)

type Handler struct {
	svc *Service
	logger *slog.Logger
}

func NewHandler(svc *Service, logger *slog.Logger) *Handler {
	return &Handler{svc: svc, logger: logger}
}


// Index home page handles GET /
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	// t := template.Must(template.ParseFiles("templates/index.html"))
	// renderTemplate(w, t, posts)
	posts, err := h.svc.ListPosts(r.Context())
	if err != nil {
		h.logger.Error("home page", "error", err)
	}
	h.logger.Debug("all posts", "posts", posts)
}
