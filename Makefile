run:
	go run main.go

build:
	CGO_ENABLED=0 go build -o bin/arsip-surat-unggulan

exec:
	./bin/arsip-surat-unggulan

start: build exec

develop:
	git push origin develop

gpo:
	git push origin

gpom:
	git push origin master

gphm:
	git push heroku master

push: gpom gphm