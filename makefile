include .env.dev
export

dev:
	go run cmd/cronjob/main.go

lint:
	golangci-lint run

dbuild:
	docker build -t rayhanhmd/orenlite-cron:latest .

dpush:
	docker push rayhanhmd/orenlite-cron:latest