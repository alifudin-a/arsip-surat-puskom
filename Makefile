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


##docker-compose##
dcup:
	docker-compose up -d

dcdown:
	docker-compose down

dcstart:
	docker-compose start

dcstop:
	docker-compose stop

dcrestart: dcdown dcup

#Dockerfile#
dbuild:
	docker build . -t arsip-surat-unggulan

dprune:
	docker image prune -f
	
logs:
	docker logs arsip-surat-unggulan -f

#Run Docker#
dredeploy: dcdown dbuild dprune dcup

#Psql#
bash:
	docker exec -it psql-arsip-surat-unggulan bash