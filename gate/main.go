package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/stellar/go/clients/horizonclient"
	_grpc "gitlab.com/velo-lab/core/gate/deposit/handler/grpc"
	_depositUcase "gitlab.com/velo-lab/core/gate/deposit/usecase"
	"gitlab.com/velo-lab/core/gate/env"
	_smartContractRepo "gitlab.com/velo-lab/core/gate/smart_contract/repository"
	_stellarRepo "gitlab.com/velo-lab/core/gate/stellar/repository"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func main() {
	env.Init()

	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zapLogger),
		)),
	)

	horizonclient := &horizonclient.Client{
		HorizonURL: "https://horizon-testnet.stellar.org/",
		HTTP:       http.DefaultClient,
	}

	ethContractBackend, err := ethclient.Dial(env.ETHNode)
	if err != nil {
		panic("unable to connect with blockchain")
	}

	sr := _stellarRepo.NewHorizonStellarRepository(horizonclient)
	smcr := _smartContractRepo.NewEthereumRepository(ethContractBackend)
	du := _depositUcase.NewDepositUsecase(sr, smcr)

	_grpc.NewDepositHandler(grpcServer, du)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", env.Port))
	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
