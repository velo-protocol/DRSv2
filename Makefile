protobuf:
	protoc -I ./proto ./proto/$(name).proto --go_out=plugins=grpc:./proto

gate-server:
	@echo "======== starting gate server ========"
	dep ensure
	docker-compose -f docker/gate/docker-compose.yaml up --build

facilitator-server:
	@echo "======== starting facilitator server ========"
	dep ensure
	docker-compose -f docker/facilitator/docker-compose.yaml up --build

contract:
	abigen -sol ./smart_contract/contracts/$(name).sol -pkg abi -out ./smart_contract/abi/$(folder)/$(name).go

format:
	go fmt ./...