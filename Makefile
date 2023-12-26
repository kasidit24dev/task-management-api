run:
	go run app/main.go

update_package:
	go mod tidy

run_test:
	go test ./...

