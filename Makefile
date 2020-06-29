main=cmd/app/main.go

run:
	@go run $(main)

dev:
	@goreload $(main)

build:
	@go build -o app $(main)