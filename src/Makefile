.DEFAULT_GOAL:=welcome_splash
RF=@go run main.go
DEPS=

app_image := "gin_processor"
port_mapping := "8086:8081"

build:
	@echo "build......"
	docker build -t ${app_image} .

run: build
	docker run -it -p ${port_mapping} ${app_image}

run:
	@go run main.go

clean:
	@echo "Cleaning up...."


