package logic_test

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/logic"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/mocks"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/crypto"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/mocks"
	"testing"
)

type helper struct {
	logic             logic.Logic
	mockDbRepo        *mocks.MockDbRepo
	mockConfiguration *mockutils.MockConfiguration
	mockController    *gomock.Controller
	logHook           *test.Hook
	done              func()
}

func initTest(t *testing.T) helper {
	mockCtrl := gomock.NewController(t)

	mockDB := mocks.NewMockDbRepo(mockCtrl)
	mockConfiguration := mockutils.NewMockConfiguration(mockCtrl)

	logger, hook := test.NewNullLogger()
	console.Logger = logger

	// to omit what loader print
	console.DefaultLoadWriter = console.Logger.Out

	return helper{
		logic:             logic.NewLogic(mockDB, mockConfiguration, nil),
		mockDbRepo:        mockDB,
		mockConfiguration: mockConfiguration,
		mockController:    mockCtrl,
		logHook:           hook,
		done: func() {
			mockCtrl.Finish()
		},
	}
}

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
	encryptedPrivateKey, nonce, _ := crypto.Encrypt([]byte("SBR25NMQRKQ4RLGNV5XB3MMQB4ADVYSMPGVBODQVJE7KPTDR6KGK3XMX"), "password")

	return entity.Account{
		PublicAddress:       "0x0f1D6Ad59AE485A9ec31b36154093820337bdEA4",
		EncryptedPrivateKey: encryptedPrivateKey,
		Nonce:               nonce,
		IsDefault:           true,
	}
}
