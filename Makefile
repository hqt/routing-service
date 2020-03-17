start-server:
	go run cmd/server/main.go
.PHONY: start-server

sample-request:
	curl --header "Content-Type: application/json" \
  	--request POST \
  	--data '{"from":"Holland Village","to":"Bugis"}' \
  	http://localhost:3000/api/simple_route
.PHONY: sample-request