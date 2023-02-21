APP_NAME?=sca

clean:
	rm -f ${APP_NAME}

build: clean
	go build -o ${APP_NAME}

test:
	go test -v -count=1 ./...

cover:
	go test -short -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
