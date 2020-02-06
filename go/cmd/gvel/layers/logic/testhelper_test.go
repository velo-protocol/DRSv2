package logic_test

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/logic"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/layers/mocks"
	"github.com/velo-protocol/DRSv2/go/cmd/gvel/utils/console"
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
		logic:             logic.NewLogic(mockDB, mockConfiguration),
		mockDbRepo:        mockDB,
		mockConfiguration: mockConfiguration,
		mockController:    mockCtrl,
		logHook:           hook,
		done: func() {
		},
	}
}
