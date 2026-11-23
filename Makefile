
.PHONY: run 
run: 
	go run .

air:
	air --tmp_dir /tmp/air --build.cmd "go build -o /tmp/air/main ." --build.bin /tmp/air/main

migrate-up:
	goose -dir migrations sqlite3 blog.db up

migrate-down:
	goose -dir migrations sqlite3 blog.db down
