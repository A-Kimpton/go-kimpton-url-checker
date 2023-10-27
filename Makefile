service_name = url-checker

build:
	go build -o bin/$(service_name)-jobdispatcher internal/cmd/jobdispatcher/main.go
	go build -o bin/$(service_name)-jobprocessor internal/cmd/jobprocessor/main.go

run: build
	bin/$(service_name)-jobdispatcher
	bin/$(service_name)-jobprocessor

run-dispatcher:
	bin/$(service_name)-jobdispatcher

run-processor:
	bin/$(service_name)-jobprocessor

docker-up:
	docker-compose -f docker/docker-compose.yaml up -d

docker-down:
	docker-compose -f docker/docker-compose.yaml down