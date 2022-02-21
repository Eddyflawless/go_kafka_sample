.DEFAULT_GOAL:=welcome_splash
RF=@go run main.go
DEPS=
#target 1
welcome_splash: 
	@echo "Gin server"

run:
	@go run main.go

all: welocome_splash run

watch: main.go 
	@echo "Reacting to file changes"
	${RF}

clean:
	@echo "Cleaning up...."


