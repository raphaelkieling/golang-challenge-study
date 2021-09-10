populate: 
	go run ./pkg/populate/populate.go
run:
	go run *.go
test:   
	go test ./...
build:
	go build
	go build ./pkg/populate/populate.go
	