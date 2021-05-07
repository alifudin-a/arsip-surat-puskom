run:
	go run main.go

build:
	go build -o arsip-surat-unggulan

exec:
	./arsip-surat-puskom

start: build exec

develop:
	git push origin develop

gpoh:
	git push origin heroku

gphm:
	git push heroku master

push: gpoh gphm