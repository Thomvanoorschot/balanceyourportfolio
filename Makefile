generate:
	rm -r ./backend/generated/proto && protoc \
		--proto_path=proto proto/*.proto \
		--go_out=backend/generated \
		--go-grpc_out=backend/generated

	rm -r ./ui/src/lib/proto && ./ui/node_modules/.bin/proto-loader-gen-types --longs=Number --enums=String --defaults --oneofs --grpcLib=@grpc/grpc-js --outDir=./ui/src/lib/proto proto/*.proto && cp proto/*.proto ./ui/src/lib/proto

up:
	docker compose up -d --build --force-recreate --remove-orphans --quiet-pull