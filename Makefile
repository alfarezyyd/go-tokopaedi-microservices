.PHONY: migrate migrate-alter migrate-create migrate-down migrate-up migrate-fresh migrate-force
ifneq (,$(wildcard services/$(SERVICE)/config.mk))
include services/$(SERVICE)/config.mk
export
endif

migrate-fresh:
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/$(DATABASE_NAME)?sslmode=disable" down
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/$(DATABASE_NAME)?sslmode=disable" up

migrate-up:
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/$(DATABASE_NAME)?sslmode=disable" up 1

migrate:
	migrate -path services/$(FEATURE)/migrations/ -database "postgres://postgres:@127.0.0.1:5432/$(DATABASE_NAME)?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/$(DATABASE_NAME)?sslmode=disable" down 1

migrate-force:
	migrate -path migrations -database "postgres://postgres:@127.0.0.1:5432/$(DATABASE_NAME)?sslmode=disable" force $(version)

migrate-create:
	migrate create -ext sql -dir services/$(FEATURE)/migrations/ create_$(name)_table


migrate-alter:
	migrate alter -ext sql -dir migrations/ alter_$(name)_table

inject:
	wire gen ./cmd/injection/injector.go

gen:
	@protoc \
	--proto_path=protobuf "protobuf/$(PROTO_NAME).proto" \
	--go_out=protobuf/genproto/ --go_opt=paths=source_relative \
  	--go-grpc_out=protobuf/genproto/ --go-grpc_opt=paths=source_relative