all: linux macos
main:
	go build -o secret cmd/main.go
run:
	go run cmd/main.go
linux:
	GOOS=linux GOARCH=amd64 go build -o secret_linux cmd/main.go
macos:
	GOOS=darwin GOARCH=arm64 go build -o secret_macos cmd/main.go
openbsd:
	GOOS=openbsd GOARCH=amd64 go build -o secret_openbsd cmd/main.go
clean:
	rm -rf secret_linux secret_macos secret_*bsd
