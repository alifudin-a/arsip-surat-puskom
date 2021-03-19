run:
	go run main.go

build:
	go build

exec:
	./arsip-surat-puskom

start: build exec

develop:
	git push origin feature/develop