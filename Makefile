start:
	#run the app
	go run ./cmd/gin/*.go

test:
	#run unit tests
	go test ./internal/tests/...