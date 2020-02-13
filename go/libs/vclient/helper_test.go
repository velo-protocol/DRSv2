package vclient

import (
	"github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/core/types"
	"github.com/Evrynetlabs/evrynet-node/crypto"
	"github.com/golang/mock/gomock"
	"github.com/velo-protocol/DRSv2/go/constants"
	"github.com/velo-protocol/DRSv2/go/libs/vclient/mocks"
	"math/big"
	"testing"
)

const (
	smartContractUrl = "http://127.0.0.1:7545"

	drsAddress            = "0x0000000000000000000000000000000000000001"
	heartAddress          = "0x0000000000000000000000000000000000000002"
	governorAddress       = "0x0000000000000000000000000000000000000003"
	trustedPartnerAddress = "0x50637DeE3598e080B7B605B00f4FfC721046E4E0"
	invalidAddress        = "GD4K"

	privateKey1 = "b673aace6739646820330920307288260703487da63525f944c96039931d8ed2"
)

type TestHelper struct {
	Client   *Client
	TxHelper *txHelper

	MockController    *gomock.Controller
	MockConnection    *mocks.MockConnection
	MockDRSContract   *mocks.MockDRSContract
	MockHeartContract *mocks.MockHeartContract
	MockTxHelper      *mocks.MockTxHelper
	Done              func()
}

func testHelperWithMock(t *testing.T) *TestHelper {
	privKey, _ := crypto.HexToECDSA(privateKey1)

	mockCtrl := gomock.NewController(t)
	mockConnection := mocks.NewMockConnection(mockCtrl)
	mockDrsContract := mocks.NewMockDRSContract(mockCtrl)
	mockHeartContract := mocks.NewMockHeartContract(mockCtrl)
	mockTxHelper := mocks.NewMockTxHelper(mockCtrl)

	client := NewClientWithOptions(&ClientOptions{
		PrivateKey:    *privKey,
		Conn:          mockConnection,
		DRSContract:   mockDrsContract,
		HeartContract: mockHeartContract,
		TxHelper:      mockTxHelper,
	})

	return &TestHelper{
		Client:            client,
		TxHelper:          NewTxHelper(mockConnection),
		MockController:    mockCtrl,
		MockConnection:    mockConnection,
		MockHeartContract: mockHeartContract,
		MockDRSContract:   mockDrsContract,
		MockTxHelper:      mockTxHelper,
		Done: func() {
			mockCtrl.Finish()
		},
	}
}

func NewMockedTx() *types.Transaction {
	return types.NewTransaction(
		99,
		common.Address{},
		big.NewInt(888),
		constants.GasLimit,
		big.NewInt(1),
		[]byte{},
	)
}
