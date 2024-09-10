.PHONY: clean build run docker-build docker-run docker-push docker-pull docker-stop k8s-deploy

clean:
	@rm -rf ./bin

build: clean
	@go build -o ./bin/go-k8s-manager ./

run: build
	@./bin/go-k8s-manager --action=serve

docker-build:
	@docker build -t rishimishra0704/go-k8s-manager .

docker-run: docker-build
	@docker run -p 8080:8080 rishimishra0704/go-k8s-manager --action=serve

docker-push: docker-build
	@docker tag rishimishra0704/go-k8s-manager:latest rishimishra0704/go-k8s-manager:latest
	@docker push rishimishra0704/go-k8s-manager:latest

docker-pull:
	@docker pull rishimishra0704/go-k8s-manager:latest

docker-stop:
	@docker stop go-k8s-manager || true

k8s-deploy: build
	@./bin/go-k8s-manager --action=k8s-deploy
