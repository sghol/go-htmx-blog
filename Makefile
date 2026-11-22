
.PHONY: run 
run: 
	go run .

migrate-up:
	goose -dir migrations sqlite3 blog.db up

migrate-down:
	goose -dir migrations sqlite3 blog.db down
