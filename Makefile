docker:=@docker
go:=@go

#######################################
# API TERRITORY COMMAND
#######################################
run:
	${go} run --race -v .

build:
	${go} build --ldflags "-r -s -w -extldflags" -o main .

#######################################
# DOCKER TERRITORY COMMAND
#######################################
dcbuild:
	${docker} build -t go-img-cropping --compress .

dcrun:
	${docker} run --name go-img-cropping -p 8080:8080 --restart=always -d go-img-cropping:latest

dcstop:
	${docker} stop go-img-cropping
	${docker} rm -f go-img-cropping