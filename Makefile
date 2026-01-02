SHELL := bash
run-orders:
# 	@go run services/orders/*.go
	@go run ./services/orders

run-kitchen:
# 	@go run services/kitchen/*.go
	@go run ./services/kitchen

gen:
	@protoc \
		--proto_path=protobuf "protobuf/orders.proto" \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/orders  \
		--go-grpc_opt=paths=source_relative

	@protoc \
		--proto_path=protobuf "protobuf/ordersv2.proto" \
		--go_out=services/common/genproto/orders/v2 --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/orders/v2  \
		--go-grpc_opt=paths=source_relative
