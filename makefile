build-image:
	docker build -t airtonlira/system_planning_financial .


run-app:
	docker-compose -f .devops/app.yaml up -d 


lint:
	golint ./...
	go fmt -n ./...

test:
	go test ./...

destroy-image:
	docker rm -f $(docker ps -a -q --filter ancestor=airtonlira/system_planning_financial)
	docker rmi airtonlira/system_planning_financial
