test:
	go test ./...

test-v:
	go test ./... -v

cover:
	go test -cover

cover-v:
	go test -cover -v

cover-html:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
