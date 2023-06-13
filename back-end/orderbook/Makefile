grpc-orderbook:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./expose/orderbook.proto

run:
	export PORT=40051 && go run *.go 

pre-build:
	go mod vendor
post-build:
	rm -rf ./vendor