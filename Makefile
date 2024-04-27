run:
	go run cmd/url-shortener/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o bin/url-shortener cmd/url-shortener/main.go

clean:
	rm -rf bin
