package logic_test

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/logic"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/mocks"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/crypto"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/mocks"
	vclientmocks "github.com/velo-protocol/DRSv2/go/libs/vclient/ivclient/mocks"
	"math/big"
	"testing"
)

type helper struct {
	logic             logic.Logic
	mockDbRepo        *mocks.MockDbRepo
	mockConfiguration *mockutils.MockConfiguration
	mockVFactory      *mocks.MockVFactoryRepo
	mockVClient       *vclientmocks.MockVClient
	mockTx            *types.Transaction
	mockTx2           *types.Transaction
	mockController    *gomock.Controller
	logHook           *test.Hook
	done              func()
}

func initTest(t *testing.T) helper {
	mockCtrl := gomock.NewController(t)

	mockDB := mocks.NewMockDbRepo(mockCtrl)
	mockVFactory := mocks.NewMockVFactoryRepo(mockCtrl)
	mockVClient := vclientmocks.NewMockVClient(mockCtrl)
	mockConfiguration := mockutils.NewMockConfiguration(mockCtrl)

	logger, hook := test.NewNullLogger()
	console.Logger = logger

	// to omit what loader print
	console.DefaultLoadWriter = console.Logger.Out

	return helper{
		logic:             logic.NewLogic(mockDB, mockConfiguration, mockVFactory),
		mockDbRepo:        mockDB,
		mockVFactory:      mockVFactory,
		mockConfiguration: mockConfiguration,
		mockVClient:       mockVClient,
		mockTx:            types.NewTransaction(1, common.Address{}, big.NewInt(1), 1, big.NewInt(1), []byte{}),
		mockTx2:           types.NewTransaction(2, common.Address{}, big.NewInt(1), 1, big.NewInt(1), []byte{}),
		mockController:    mockCtrl,
		logHook:           hook,
		done: func() {
			mockCtrl.Finish()
		},
	}
}

var (
	pub  = "0xf41E18a9573832265F74a671d3E275ec76790b5C"
	priv = "6d71af6c908ff8b618825926f1004431915faf9b66238c30af9f86438d2bcd89"
)

func arrayOfAccountsBytes() [][]byte {
	return [][]byte{
		accountsBytes(),
	}
}

func accountsBytes() []byte {
	accountBytes, _ := json.Marshal(accountEntity())
	return accountBytes
}

func accountEntity() entity.Account {
	encryptedPrivateKey, nonce, _ := crypto.Encrypt([]byte("6d71af6c908ff8b618825926f1004431915faf9b66238c30af9f86438d2bcd89"), "password")

	return entity.Account{
		PublicAddress:       "0xf41E18a9573832265F74a671d3E275ec76790b5C",
		EncryptedPrivateKey: encryptedPrivateKey,
		Nonce:               nonce,
		IsDefault:           true,
	}
}
