BLUE = \033[34m
NC = \033[0m

help:  ## Show help message.
	@printf "Usage:\n"
	@printf "  make $(BLUE)<target>$(NC)\n\n"
	@printf "Targets:\n"
	@perl -nle'print $& if m{^[a-zA-Z0-9_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; \
		{printf "$(BLUE)  %-18s$(NC) %s\n", $$1, $$2}'


install: ## Install the dependencies
	./script/run.sh install


serve: generate ## Serve the application with hot reload in dev mode
	./script/run.sh dev serve

generate: ## Generate the gorm queries
	./script/run.sh generate

dev-up: ## Start the container for development
	./script/run.sh dev start


dev-migrate: ## Run the migrations
	./script/run.sh dev migrate


dev-down: ## Down the container for development
	./script/run.sh dev stop


dev-teardown: ## Down the container and release all resources
	./script/run.sh dev teardown

stage-up: ## Start the container for stage environment
	./script/run.sh stage start --build

test: ## Run the tests
	go test -v -cover ./...