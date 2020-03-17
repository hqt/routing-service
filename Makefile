ci-test:
	docker-compose -f docker/docker-compose-test.yml up \
    --force-recreate \
    --abort-on-container-exit \
    --exit-code-from app \
    --build

start-server-native:
	go run cmd/server/main.go
.PHONY: start-server-native

start-server:
	docker-compose -f docker/docker-compose.yml up --build
.PHONY: start-server

simple-request:
	curl --header "Content-Type: application/json" \
  	--request POST \
  	--data '{"from":"Holland Village","to":"Bugis"}' \
  	http://0.0.0.0:3000/api/simple_route
.PHONY: simple-request

advanced-request:
	curl --header "Content-Type: application/json" \
  	--request POST \
  	--data '{"from":"Boon Lay","to":"Little India", "start_time": "2019-01-31T16:00"}' \
  	http://0.0.0.0:3000/api/advanced_route
