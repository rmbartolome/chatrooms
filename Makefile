BUILDPATH=$(CURDIR)
API_NAME=chatrooms-master-mold

build: 
	@echo "Creando Binario ..."
	@go build -mod=vendor -ldflags '-s -w' -o $(BUILDPATH)/build/bin/${API_NAME} cmd/server/main.go
	@echo "Binario generado en build/bin/${API_NAME}"

test: 
	@echo "Ejecutando tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

coverage: 
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -html=coverfile_out -o coverfile_out.html

docker:
	@docker build . -t customer-debt-assesment:latest -f iaas/Dockerfile

.PHONY: test build
