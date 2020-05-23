
build-client:
	cd client; npm run build

statik:
	statik -src=./client/public -dest=./internal/pkg

build: build-client statik
	go build -o draught-log main.go

dev:
	cd client; npm run dev