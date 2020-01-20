abigen/install:
	go get github.com/ethereum/go-ethereum
	cd $(GOPATH)/src/github.com/ethereum/go-ethereum
	godep go install ./cmd/abigen

abigen/gen/all:
	yarn --cwd ./.. run abi:extract
	abigen --abi ../build/contracts/DigitalReserveSystem.abi --pkg vabi --type DigitalReserveSystem --out abi/digital_reserve_system.go
	abigen --abi ../build/contracts/Heart.abi --pkg vabi --type Heart --out abi/heart.go

