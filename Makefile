gorun:
	go run main.go

#gobuild:
#CGO_ENABLED=0 go build -o bin/arsip-surat-unggulan
#exec:
#./bin/arsip-surat-unggulan

gobuild:
	go build -o arsip-surat-unggulan

exec:
	./arsip-surat-unggulan

startapp: gobuild exec

develop:
	git push origin develop

gpo:
	git push origin

gpom:
	git push origin master

gphm:
	git push heroku master

push: gpom gphm

##server##
appname := arsip_surat

redeploy: build restart log

run:
	go run main.go
build:
	go build -o $(appname) main.go && chmod +x $(appname)
update:
	supervisorctl update
restart:
	supervisorctl restart $(appname)
start:
	supervisorctl start $(appname)
stop:
	supervisorctl stop $(appname)
createlog:
	@touch /var/log/$(appname).log
log: 
	tail -f /var/log/$(appname).log