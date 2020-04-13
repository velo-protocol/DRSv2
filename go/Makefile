abigen/install:
	go get github.com/ethereum/go-ethereum
	cd $(GOPATH)/src/github.com/ethereum/go-ethereum
	godep go install ./cmd/abigen

abigen/gen/all:
	yarn --cwd ./.. run abi:extract
	abigen --abi ../build/contracts/abi/DigitalReserveSystem.abi --bin ../build/contracts/bin/DigitalReserveSystem.bin --pkg vabi --type DigitalReserveSystem --out abi/digital_reserve_system.go
	abigen --abi ../build/contracts/abi/Heart.abi --bin ../build/contracts/bin/Heart.bin --pkg vabi --type Heart --out abi/heart.go
	abigen --abi ../build/contracts/abi/StableCredit.abi --bin ../build/contracts/bin/StableCredit.bin --pkg vabi --type StableCredit --out abi/stable_credit.go
	abigen --abi ../build/contracts/abi/Token.abi --bin ../build/contracts/bin/Token.bin --pkg vabi --type Token --out abi/token.go
	abigen --abi ../build/contracts/abi/PriceFeeders.abi --bin ../build/contracts/bin/PriceFeeders.bin --pkg vabi --type PriceFeeders --out abi/price_feeders.go
	abigen --abi ../build/contracts/abi/Feeder.abi --bin ../build/contracts/bin/Feeder.bin --pkg vabi --type Feeder --out abi/feeder.go
	abigen --abi ../build/contracts/abi/Medianizer.abi --bin ../build/contracts/bin/Medianizer.bin --pkg vabi --type Medianizer --out abi/medianizer.go
	abigen --abi ../build/contracts/abi/Lag.abi --bin ../build/contracts/bin/Lag.bin --pkg vabi --type Lag --out abi/lag.go
	abigen --abi ../build/contracts/abi/Price.abi --bin ../build/contracts/bin/Price.bin --pkg vabi --type Price --out abi/price.go

format:
	go fmt ./...

ci_test:
	go test --tags="unit" ./... -v -coverprofile .coverage.txt

mockgen/all:
	mockgen \
		-source=./libs/vclient/interface.go \
		-destination=./libs/vclient/mocks/mocked_interfaces.go \
		-package mocks \
		-imports vabi=github.com/velo-protocol/DRSv2/go/abi \
		-mock_names Connection=MockConnection,DRSContract=MockDRSContract,HeartContract=MockHeartContract,TxHelper=MockTxHelper
	mockgen \
		-source=./libs/vclient/ivclient/vclient.go \
		-destination=./libs/vclient/ivclient/mocks/mocked_interfaces.go \
		-package mocks \
		-imports vabi=github.com/velo-protocol/DRSv2/go/abi \
		-mock_names VClient=MockVClient
