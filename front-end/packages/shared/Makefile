# grpc-orderbook:
# 	rm -rf ./grpc/*
# 	protoc -I=../../../nftm/apis/grpc/expose/orderbook orderbook.proto \
#     --js_out=import_style=commonjs,binary:grpc/ \
#     --grpc-web_out=import_style=typescript,mode=grpcwebtext:./grpc/

grpc-orderbook:
	rm -rf ./grpc/*
	npx protoc --ts_out ./grpc --proto_path ../../../nftm/apis/grpc/expose/orderbook orderbook.proto

graphql-types:
	apollo client:download-schema --endpoint=https://jkilkuhllzforn3ut2mbadpcoy.appsync-api.eu-west-1.amazonaws.com/graphql  graphql-schema.json --header="x-api-key:da2-kujwfmlxdnd3ffday7pj653vfi"  
	apollo codegen:generate --localSchemaFile=graphql-schema.json --target=typescript --tagName=gql  --includes '**/*.ts' 
