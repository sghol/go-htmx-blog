package main

import (
	"database/sql"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	_ "modernc.org/sqlite"
	// "github.com/sghol/go-htmx-blog/db"
	"github.com/sghol/go-htmx-blog/internal/post"
)

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
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
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

func OpenDB(DBName string) (*sql.DB, error) {
	sqlDB, err := sql.Open("sqlite", DBName)
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return nil, err
	}
	return sqlDB, nil
}

func main() {
	// setup logger
	logger := setupLogger()

	// get blog db
	blogDB, err := OpenDB("blog.db")
	if err != nil {
		logger.Error("database error", "error", err)
		os.Exit(1)
	}

	// requests multiplexer aka url router
	mux := http.NewServeMux()

	// queries := db.New(blogDB)

	// post package
	postSvc := post.NewService(blogDB)
	postHandler := post.NewHandler(postSvc, logger)


	// handlers
	mux.HandleFunc("GET /{$}", postHandler.Index)
	mux.HandleFunc("GET /static/htmx.min.js", ServeHTMX)
	mux.HandleFunc("GET /static/", ServeStaticFiles)

	// wrap the mux with logger middleware
	handler := requestLogger(mux)

	// server
	log.Println("running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
