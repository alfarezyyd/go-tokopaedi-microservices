.PHONY: migrate-fresh
.PHONY: migrate-up
.PHONY: migrate-down
PROTO_NAME=$(proto)

migrate-fresh:
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/go_ufis_api?sslmode=disable" down
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/go_ufis_api?sslmode=disable" up

migrate-up:
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/go_ufis_api?sslmode=disable" up 1

migrate:
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/go_ufis_api?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/go_ufis_api?sslmode=disable" down 1

migrate-force:
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/go_ufis_api?sslmode=disable" force $(version)

migrate-create:
	migrate create -ext sql -dir migrations/ create_$(name)_table

migrate-alter:
	migrate alter -ext sql -dir migrations/ alter_$(name)_table

inject:
	wire gen ./cmd/injection/injector.go

gen:
	@protoc \
	--proto_path=protobuf "protobuf/$(PROTO_NAME).proto" \
	--go_out=protobuf/genproto/ --go_opt=paths=source_relative \
  	--go-grpc_out=protobuf/genproto/ --go-grpc_opt=paths=source_relative