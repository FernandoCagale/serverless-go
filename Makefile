compile:
	GOOS=linux go build -o bin/main functions/main.go
.PHONY: compile