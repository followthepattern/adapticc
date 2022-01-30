dev-adapticc-run:
	export GO111MODULE=on && go run ./cmd/adapticc

docker-adapticc-build:
	docker build -f ./build/docker/cms/Dockerfile -t cms .

hot-run-adapticc:
	CompileDaemon -log-prefix=false -build="go build -o ./adapticc ./cmd/adapticc" -command="./adapticc" -color=true

lint:
	golangci-lint run --config .golangci.yml ./...

test:
	go test ./...