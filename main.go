package main

import (
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type Post struct {
	ID          int
	Title       string
	Status      string
	Body        string
	IsPublished bool
}

var posts = []Post{
	{ID: 1, Title: "Web programin with go lang", Body: "body of the post", IsPublished: true},
	{ID: 1, Title: "Python Django", Body: "second post", IsPublished: true},
}

// renderTemplate render template with data
func renderTemplate(w http.ResponseWriter, t *template.Template, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Serve htmx to the browser with correct header metadata
// handles /static/htmx.min.js
func ServeHTMX(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, "./static/htmx.min.js.gz")
}

// ServeStaticFiles css, images, assets, ...
func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	fs.ServeHTTP(w, r)
}

// setup a logger
func setupLogger() *slog.Logger {
	var handler slog.Handler
	handler = slog.NewTextHandler(os.Stderr, nil)

	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}

// logger middleware
func requestLogger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), "request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

// Index home page
// handles GET /
func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	t := template.Must(template.ParseFiles("templates/index.html"))
	renderTemplate(w, t, posts)
}


func main() {
	// requests multiplexer aka url router
	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("GET /", Index)
	mux.HandleFunc("GET /static/htmx.min.js", ServeHTMX)
	mux.HandleFunc("GET /static/", ServeStaticFiles)

	// wrap the mux with logger middleware
	handler := requestLogger(mux)

	// server
	log.Println("running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
