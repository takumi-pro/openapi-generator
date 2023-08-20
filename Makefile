.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: destroy
destroy:
	docker-compose down --rmi all --volumes --remove-orphans

.PHONY: openapi-generate
openapi-generate:
	openapi-generator generate -i reference/todos.yaml -g go-server -o ./