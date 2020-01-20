abigen/install:
	go get github.com/ethereum/go-ethereum
	cd $(GOPATH)/src/github.com/ethereum/go-ethereum
	godep go install ./cmd/abigen

abigen/gen/all:
	yarn --cwd ./.. run abi:extract
	abigen --abi ../build/contracts/abi/DigitalReserveSystem.abi --bin ../build/contracts/bin/DigitalReserveSystem.bin --pkg vabi --type DigitalReserveSystem --out abi/digital_reserve_system.go
	abigen --abi ../build/contracts/abi/Heart.abi --bin ../build/contracts/bin/Heart.bin --pkg vabi --type Heart --out abi/heart.go

