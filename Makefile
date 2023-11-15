build:
	sh ./deployments/build_app.sh
build-postgres:
	sh ./deployments/build_postgres.sh
build-nginx:
	sh ./deployments/build_nginx.sh

lint:
	golangci-lint run
lint-all:
	golangci-lint run --enable-all
lint-to-file:
	golangci-lint run --out-format html > lint.html
lint-all-to-file:
	golangci-lint run --enable-all --out-format html > lint.html

test:
	go test -v count=10 ./...

test_coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out
