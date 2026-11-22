package post

import (
	"github.com/sghol/go-htmx-blog/response"
	"log/slog"
	"net/http"
)

type Handler struct {
	svc    *Service
	logger *slog.Logger
	resp      *response.HTMLRenderer
}

func NewHandler(svc *Service, logger *slog.Logger, resp *response.HTMLRenderer) *Handler {
	return &Handler{svc: svc, logger: logger, resp: resp}
}

// Index home page handles GET /
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	posts, err := h.svc.ListPosts(r.Context())
	if err != nil {
		h.resp.ServerError(w, err)
		return
	}
	h.resp.RenderHTML(w, "index.html", posts)
}
