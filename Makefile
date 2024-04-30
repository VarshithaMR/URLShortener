#Name of the binary executable
NAME=url-shortner-application
#Local environment
LOCAL_ENV=Path="$(GOPATH)/bin:$(PATH)"
#Set Go module
GO=GO111MODULE=on $(LOCAL_ENV) go
#Root files
ROOT_FILE=main/main.go

#set c-go, OS and Arch
CGO=0
GOOS=linux
GOARCH=amd64

build:
	@echo "Building the service...."
	$(GO) mod tidy
	CGO_ENABLED=${CGO} GOOS=${GOOS} GOARCH=${GOARCH} $(GO) build -o $(NAME) $(ROOT_FILE)
.PHONY: build

run:
	@echo"Running the application...."
	$(GO) run $(ROOT_FILE)