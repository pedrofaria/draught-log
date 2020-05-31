
build-client:
	cd client; npm run build

statik:
	statik -src=./client/public -dest=./internal/pkg

build: build-client statik
	go build -o draught-log main.go

dev:
	cd client; npm run dev

deps:
	cd client; npm install
	go get github.com/rakyll/statik@v0.1.7
	go mod download

docker:
	docker build -t pedrofaria/draught-log:$${TAG:-latest} -f build/docker/Dockerfile .
