.SILENT:
.PHONY:

PATH_TO_PROTO=./api/proto
PROTO_OUT=./pkg/api

proto:
	protoc \
	--go_out=$(PROTO_OUT) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_OUT) \
	--go-grpc_opt=paths=source_relative \
	--proto_path=$(PATH_TO_PROTO) $(PATH_TO_PROTO)/*.proto