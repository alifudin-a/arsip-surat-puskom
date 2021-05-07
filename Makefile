run:
	go run main.go

build:
	go build -o bin/arsip-surat-unggulan

exec:
	./bin/arsip-surat-unggulan

start: build exec

develop:
	git push origin

gpo:
	git push origin heroku

gphm:
	git push heroku master

push: gpoh gphm