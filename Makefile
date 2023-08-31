generate:
	@rm -rf ./src/swaggerr/generated
	@mkdir -p src/swagger/generated
	@swagger generate server \
		-f ./swagger.yaml \
		-t ./src/swagger/generated -C ./templates/default.yml \
		--template-dir ./templates/tmpl \
		--name trainee-assignment

up:
	@docker-compose -f docker-compose.dev.yml up --build -d

down:
	@docker-compose -f docker-compose.dev.yml down

dry-run:
	@go run main.go

test:
	@docker-compose -f docker-compose.test.yml down
	@docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from tests --quiet-pull
