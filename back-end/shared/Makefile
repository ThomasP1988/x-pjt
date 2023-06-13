grpc-orderbook:
	protoc --go_out=./orderbook/grpc --go_opt=paths=source_relative \
    --go-grpc_out=./orderbook/grpc --go-grpc_opt=paths=source_relative \
   	-I=./orderbook/grpc orderbook.proto

pre-build:
	go mod vendor
post-build:
	rm -rf ./vendor