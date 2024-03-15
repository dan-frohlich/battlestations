# Change these variables as necessary.
CLI_PACKAGE_PATH := ./cmd/cli
BINARY_NAME := bs-printer

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## check: run quality control checks
.PHONY: check
check:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## build: build the application
.PHONY: build
build:
	# Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -o ./bin/${BINARY_NAME} ${CLI_PACKAGE_PATH}


# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## launch: lauch in fly
# .PHONY: launch
# launch: confirm tidy check no-dirty
# 	~/.fly/bin/fly launch

## deploy: deploy the application to fly.io
# .PHONY: deploy
# deploy: confirm tidy check no-dirty
# 	~/.fly/bin/flyctl deploy

## push: push changes to the remote Git repository
.PHONY: push
push: tidy check no-dirty
	git push

# production/deploy: deploy the application to production
#.PHONY: production/deploy
#production/deploy: confirm tidy check no-dirty
#	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/${BINARY_NAME}_linux_amd64_web ${WEB_PACKAGE_PATH}
#	upx -5 ./bin/${BINARY_NAME}_linux_amd64_web
#	# Include additional deployment steps here...
#	
