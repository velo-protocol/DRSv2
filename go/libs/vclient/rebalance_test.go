package vclient

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/abi"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"math/big"
	"testing"
)

func TestRebalance(t *testing.T) {
	t.Run("success, rebalance with stableCredit 1 loop", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		stableCreditSize := uint8(1)
		stableCreditAddress := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"
		collateralAssetCode := utils.StringToByte32("VELO")

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditSize, nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditAddress, nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCode, &stableCreditId, nil)

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Rebalance(string,bytes32,uint256,uint256)")),
						},
					},
				},
			}, nil)

		testHelper.MockTxHelper.EXPECT().
			ExtractRebalanceEvent("Rebalance", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemRebalance{
				AssetCode:           stableCreditAssetCode,
				CollateralAssetCode: collateralAssetCode,
				PresentAmount:       big.NewInt(1000000000),
				RequiredAmount:      big.NewInt(1000000000),
			}, nil)

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.Equal(t, "100.0000000", result.RebalanceTransactions[0].RequiredAmount)
		assert.Equal(t, "100.0000000", result.RebalanceTransactions[0].PresentAmount)
		assert.Equal(t, stableCreditAssetCode, result.RebalanceTransactions[0].AssetCode)
		assert.Equal(t, utils.Byte32ToString(collateralAssetCode), result.RebalanceTransactions[0].CollateralAssetCode)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("success, rebalance with stableCredit 2 loop", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		stableCreditSize := uint8(2)
		stableCreditAddress := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		nextStableCreditAddress := common.HexToAddress("0x8F384Dd497A06941531ED277Cc0Cd34e7d594323")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCodeVUSD := "vUSD"
		stableCreditAssetCodeVTHB := "vTHB"
		collateralAssetCode := utils.StringToByte32("VELO")

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditSize, nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditAddress, nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCodeVUSD, &stableCreditId, nil)
		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCodeVTHB, &stableCreditId, nil)

		testHelper.MockHeartContract.EXPECT().
			GetNextStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{}), gomock.AssignableToTypeOf(stableCreditId)).
			Return(nextStableCreditAddress, nil)

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCodeVUSD).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Rebalance(string,bytes32,uint256,uint256)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractRebalanceEvent("Rebalance", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemRebalance{
				AssetCode:           stableCreditAssetCodeVUSD,
				CollateralAssetCode: collateralAssetCode,
				PresentAmount:       big.NewInt(1000000000),
				RequiredAmount:      big.NewInt(1000000000),
			}, nil)

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCodeVTHB).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Rebalance(string,bytes32,uint256,uint256)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractRebalanceEvent("Rebalance", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemRebalance{
				AssetCode:           stableCreditAssetCodeVTHB,
				CollateralAssetCode: collateralAssetCode,
				PresentAmount:       big.NewInt(1000000000),
				RequiredAmount:      big.NewInt(1000000000),
			}, nil)

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "100.0000000", result.RebalanceTransactions[0].RequiredAmount)
		assert.Equal(t, "100.0000000", result.RebalanceTransactions[0].PresentAmount)
		assert.Equal(t, stableCreditAssetCodeVUSD, result.RebalanceTransactions[0].AssetCode)

		assert.Equal(t, "100.0000000", result.RebalanceTransactions[1].RequiredAmount)
		assert.Equal(t, "100.0000000", result.RebalanceTransactions[1].PresentAmount)
		assert.Equal(t, stableCreditAssetCodeVTHB, result.RebalanceTransactions[1].AssetCode)
	})

	t.Run("success, rebalance with stableCredit 3 loop", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		stableCreditSize := uint8(3)
		stableCreditAddress := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		nextStableCreditAddress := common.HexToAddress("0x8F384Dd497A06941531ED277Cc0Cd34e7d594323")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCodeVUSD := "vUSD"
		stableCreditAssetCodeVTHB := "vTHB"
		stableCreditAssetCodeVSGD := "vSGD"

		collateralAssetCodeVELO := utils.StringToByte32("VELO")
		collateralAssetCodeETH := utils.StringToByte32("ETH")

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditSize, nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditAddress, nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCodeVUSD, &stableCreditId, nil)
		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCodeVTHB, &stableCreditId, nil)
		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCodeVSGD, &stableCreditId, nil)

		testHelper.MockHeartContract.EXPECT().
			GetNextStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{}), gomock.AssignableToTypeOf(stableCreditId)).
			Return(nextStableCreditAddress, nil)
		testHelper.MockHeartContract.EXPECT().
			GetNextStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{}), gomock.AssignableToTypeOf(stableCreditId)).
			Return(nextStableCreditAddress, nil)

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCodeVUSD).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Rebalance(string,bytes32,uint256,uint256)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractRebalanceEvent("Rebalance", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemRebalance{
				AssetCode:           stableCreditAssetCodeVUSD,
				CollateralAssetCode: collateralAssetCodeVELO,
				PresentAmount:       big.NewInt(1000000000),
				RequiredAmount:      big.NewInt(1000000000),
			}, nil)

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCodeVTHB).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Rebalance(string,bytes32,uint256,uint256)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractRebalanceEvent("Rebalance", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemRebalance{
				AssetCode:           stableCreditAssetCodeVTHB,
				CollateralAssetCode: collateralAssetCodeETH,
				PresentAmount:       big.NewInt(1000000000),
				RequiredAmount:      big.NewInt(1000000000),
			}, nil)

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCodeVSGD).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Rebalance(string,bytes32,uint256,uint256)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractRebalanceEvent("Rebalance", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemRebalance{
				AssetCode:           stableCreditAssetCodeVSGD,
				CollateralAssetCode: collateralAssetCodeVELO,
				PresentAmount:       big.NewInt(1000000000),
				RequiredAmount:      big.NewInt(1000000000),
			}, nil)

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("success, no event emit rebalance equilibrium", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		stableCreditSize := uint8(1)
		stableCreditAddress := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditSize, nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditAddress, nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCode, &stableCreditId, nil)

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{},
				},
			}, nil)

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.NotNil(t, result)
		assert.NoError(t, err)
	})

	t.Run("fail, call Heart Contract GetStableCreditCount", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(0), errors.New("error here"))

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call Heart Contract GetStableCreditCount with stableCreditSize 0", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(0), nil)

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "stable credit not found", err.Error())
	})

	t.Run("fail, call Heart Contract GetRecentStableCredit", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(1), nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(common.Address{}, errors.New("error here"))

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call txHelper StableCreditAssetCode", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(1), nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"), nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"))).
			Return(nil, nil, errors.New("error here"))

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call Drs Contract Rebalance 2 loop but got error from call Heart Contract GetNextStableCredit", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"
		collateralAssetCode := utils.StringToByte32("VELO")

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(2), nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"), nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"))).
			Return(&stableCreditAssetCode, &stableCreditId, nil)

		testHelper.MockHeartContract.EXPECT().
			GetNextStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{}), gomock.AssignableToTypeOf(stableCreditId)).
			Return(common.Address{}, errors.New("error here"))

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Rebalance(string,bytes32,uint256,uint256)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractRebalanceEvent("Rebalance", gomock.AssignableToTypeOf(&types.Log{})).
			Return(&vabi.DigitalReserveSystemRebalance{
				AssetCode:           stableCreditAssetCode,
				CollateralAssetCode: collateralAssetCode,
				PresentAmount:       big.NewInt(1000000000),
				RequiredAmount:      big.NewInt(1000000000),
			}, nil)

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call Drs Contract Rebalance throw error", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(2), nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"), nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"))).
			Return(&stableCreditAssetCode, &stableCreditId, nil)

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCode).
			Return(nil, errors.New("error here"))

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call txHelper.ConfirmTx", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(2), nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"), nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"))).
			Return(&stableCreditAssetCode, &stableCreditId, nil)

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(nil, errors.New("error here"))

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call txHelper.ExtractRebalanceEvent", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		stableCreditSize := uint8(1)
		stableCreditAddress := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditSize, nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditAddress, nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCode, &stableCreditId, nil)

		testHelper.MockDRSContract.EXPECT().
			Rebalance(gomock.AssignableToTypeOf(&bind.TransactOpts{}), stableCreditAssetCode).
			Return(&types.Transaction{}, nil)
		testHelper.MockTxHelper.EXPECT().
			ConfirmTx(gomock.AssignableToTypeOf(context.Background()), gomock.AssignableToTypeOf(&types.Transaction{}), gomock.AssignableToTypeOf(common.Address{})).
			Return(&types.Receipt{
				Logs: []*types.Log{
					{
						Topics: []common.Hash{
							crypto.Keccak256Hash([]byte("Rebalance(string,bytes32,uint256,uint256)")),
						},
					},
				},
			}, nil)
		testHelper.MockTxHelper.EXPECT().
			ExtractRebalanceEvent("Rebalance", gomock.AssignableToTypeOf(&types.Log{})).
			Return(nil, errors.New("error here"))

		result, err := testHelper.Client.Rebalance(context.Background(), &RebalanceInput{})

		assert.Nil(t, result)
		assert.Error(t, err)
	})
}
