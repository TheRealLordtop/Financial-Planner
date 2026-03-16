APP_NAME=app

build:
	go build -o ./bin/$(APP_NAME) ./cmd/app

run:
	go run ./cmd/app

clean:
	rm -rf ./bin