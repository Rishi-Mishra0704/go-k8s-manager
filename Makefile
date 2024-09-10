clean:
	@rm -rf ./bin

build: clean
	@go build -o ./bin/go-k8s-manager ./

run: build
	@./bin/go-k8s-manager