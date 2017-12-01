build:
	env GOOS=windows env GOARCH=amd64 go build -o windows/main.exe main.go
	go build -o mac/main main.go

install:
	go get github.com/gin-gonic/gin

run:
	go run main.go
