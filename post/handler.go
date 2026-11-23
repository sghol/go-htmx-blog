package post

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/sghol/go-htmx-blog/response"
)

type Handler struct {
	svc    *Service
	logger *slog.Logger
	resp   *response.HTMLRenderer
}

func NewHandler(svc *Service, logger *slog.Logger, resp *response.HTMLRenderer) *Handler {
	return &Handler{svc: svc, logger: logger, resp: resp}
}

// Index home page handles GET /
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	posts, err := h.svc.ListPosts(r.Context())
	if err != nil {
		h.resp.ServerError(w, err)
		return
	}
	h.resp.RenderHTML(w, "home.html", posts)
}

// CreatePost handles creating new post by POST /posts
func (h *Handler) RenderCreatePostForm(w http.ResponseWriter, r *http.Request) {
	h.resp.RenderHTML(w, "post_create.html", nil)
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	_, err := h.svc.CreatePost(r.Context(), CreatePostInput{
		Title:       strings.TrimSpace(r.FormValue("title")),
		Body:        r.FormValue("body"),
		IsPublished: r.FormValue("is_published") != "",
	})
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Error("Create new post", "error", err)
		return
	}
}
