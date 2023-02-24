dockerBuild:
	docker-compose up --build

dockerUp:
	docker-compose up

dockerDown:
	docker-compose down

dockerLog:
	docker-compose logs

goBuildR:
	go build cmd/2FA-with-golang/main.go && ./main