all: 
	go build -o build/TapTake-server

debug:
	go build -o build/TapTake-server
	ENV_MODE=debug PORT=9090 ./build/TapTake-server
	
run:
	go build -o build/TapTake-server
	./build/TapTake-server
