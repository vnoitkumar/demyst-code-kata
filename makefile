format:
	go fmt ./...

test:
	go test ./... -short -count=1 -p=1

test_with_logs:
	go test ./... -short -count=1 -p=1 -v

test_with_coverage:
	go test -coverprofile cover.out ./... -short -count=1 -p=1

open_coverage:
	go tool cover -html cover.out -o cover.html && open cover.html

pre_commit:
	go mod tidy
	go mod vendor
	go vet ./...
	go fmt ./...
	golangci-lint run -v ./... --timeout=180s
	make build

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/macos-intel-amd64 . && GOOS=darwin GOARCH=arm64 go build -o bin/macos-apple-silicon-arm64 . && GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64.exe . && GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64 . 

dockerize:
	docker build --rm -t demyst-code-kata:lastest -f infrastructure/Dockerfile .

stop_container:
	docker stop demyst-code-kata-container || true

remove_container:
	docker rm demyst-code-kata-container || true

run:
	docker run -d --name demyst-code-kata-container demyst-code-kata:lastest

logs:
	docker logs -ft demyst-code-kata-container

start:
	make dockerize && make stop_container && make remove_container && make run && make logs


