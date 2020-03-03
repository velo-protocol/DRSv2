package vclient

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"testing"
)

func TestCollateralHealthCheck(t *testing.T) {
	t.Run("success, 1 stableCredit loop", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		requiredAmount, _ := utils.StringToAmount("100")
		presentAmount, _ := utils.StringToAmount("200")
		collateralAssetCode := utils.StringToByte32("VELO")
		collateralAssetAddress := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f1")

		stableCreditSize := uint8(1)
		stableCreditAddress := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"

		input := &CollateralHealthCheckInput{}
		toAbiInput := &CollateralHealthCheckAbiInput{
			AssetCode: stableCreditAssetCode,
		}

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
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(collateralAssetAddress, collateralAssetCode, requiredAmount, presentAmount, nil)

		result, err := testHelper.Client.CollateralHealthCheck(input)
		requiredAmountVELO := result[0].RequiredAmount
		presentAmountVELO := result[0].PresentAmount

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "VELO", result[0].CollateralAssetCode)
		assert.Equal(t, collateralAssetAddress.String(), result[0].CollateralAssetAddress)
		assert.Equal(t, "100.0000000", requiredAmountVELO)
		assert.Equal(t, "200.0000000", presentAmountVELO)
	})

	t.Run("success, 2 stableCredit loop", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		requiredAmount, _ := utils.StringToAmount("100")
		presentAmount, _ := utils.StringToAmount("200")
		collateralAssetCodeVELO := utils.StringToByte32("VELO")
		collateralAssetAddressVELO := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f1")
		collateralAssetCodeETH := utils.StringToByte32("ETH")
		collateralAssetAddressETH := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f2")

		stableCreditSize := uint8(2)
		stableCreditAddress := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		nextStableCreditAddress := common.HexToAddress("0x8F384Dd497A06941531ED277Cc0Cd34e7d594323")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"

		input := &CollateralHealthCheckInput{}
		toAbiInput := &CollateralHealthCheckAbiInput{
			AssetCode: stableCreditAssetCode,
		}

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditSize, nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditAddress, nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCode, &stableCreditId, nil)
		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCode, &stableCreditId, nil)

		testHelper.MockHeartContract.EXPECT().
			GetNextStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{}), gomock.AssignableToTypeOf(stableCreditId)).
			Return(nextStableCreditAddress, nil)

		testHelper.MockDRSContract.EXPECT().
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(collateralAssetAddressVELO, collateralAssetCodeVELO, requiredAmount, presentAmount, nil)
		testHelper.MockDRSContract.EXPECT().
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(collateralAssetAddressETH, collateralAssetCodeETH, requiredAmount, presentAmount, nil)

		result, err := testHelper.Client.CollateralHealthCheck(input)
		requiredAmountVELO := result[0].RequiredAmount
		presentAmountVELO := result[0].PresentAmount
		requiredAmountETH := result[1].RequiredAmount
		presentAmountETH := result[1].PresentAmount

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "VELO", result[0].CollateralAssetCode)
		assert.Equal(t, collateralAssetAddressVELO.String(), result[0].CollateralAssetAddress)
		assert.Equal(t, "100.0000000", requiredAmountVELO)
		assert.Equal(t, "200.0000000", presentAmountVELO)
		assert.Equal(t, "ETH", result[1].CollateralAssetCode)
		assert.Equal(t, collateralAssetAddressETH.String(), result[1].CollateralAssetAddress)
		assert.Equal(t, "100.0000000", requiredAmountETH)
		assert.Equal(t, "200.0000000", presentAmountETH)
	})

	t.Run("success, 3 stableCredit in 2 collateralAssetCode loop", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		requiredAmountVELO, _ := utils.StringToAmount("100")
		requiredAmountETH, _ := utils.StringToAmount("100")
		presentAmountVELO, _ := utils.StringToAmount("200")
		presentAmountETH, _ := utils.StringToAmount("300")
		collateralAssetCodeVELO := utils.StringToByte32("VELO")
		collateralAssetAddressVELO := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f1")
		collateralAssetCodeETH := utils.StringToByte32("ETH")
		collateralAssetAddressETH := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f2")

		stableCreditSize := uint8(3)
		stableCreditAddress := common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		nextStableCreditAddress := common.HexToAddress("0x8F384Dd497A06941531ED277Cc0Cd34e7d594323")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"

		input := &CollateralHealthCheckInput{}
		toAbiInput := &CollateralHealthCheckAbiInput{
			AssetCode: stableCreditAssetCode,
		}

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditSize, nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(stableCreditAddress, nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCode, &stableCreditId, nil)
		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCode, &stableCreditId, nil)
		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(stableCreditAddress)).
			Return(&stableCreditAssetCode, &stableCreditId, nil)

		testHelper.MockHeartContract.EXPECT().
			GetNextStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{}), gomock.AssignableToTypeOf(stableCreditId)).
			Return(nextStableCreditAddress, nil)
		testHelper.MockHeartContract.EXPECT().
			GetNextStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{}), gomock.AssignableToTypeOf(stableCreditId)).
			Return(nextStableCreditAddress, nil)

		testHelper.MockDRSContract.EXPECT().
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(collateralAssetAddressVELO, collateralAssetCodeVELO, requiredAmountVELO, presentAmountVELO, nil)
		testHelper.MockDRSContract.EXPECT().
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(collateralAssetAddressVELO, collateralAssetCodeVELO, requiredAmountVELO, presentAmountVELO, nil)
		testHelper.MockDRSContract.EXPECT().
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(collateralAssetAddressETH, collateralAssetCodeETH, requiredAmountETH, presentAmountETH, nil)

		result, err := testHelper.Client.CollateralHealthCheck(input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "VELO", result[0].CollateralAssetCode)
		assert.Equal(t, collateralAssetAddressVELO.String(), result[0].CollateralAssetAddress)
		assert.Equal(t, "200.0000000", result[0].RequiredAmount)
		assert.Equal(t, "400.0000000", result[0].PresentAmount)
		assert.Equal(t, "ETH", result[1].CollateralAssetCode)
		assert.Equal(t, collateralAssetAddressETH.String(), result[1].CollateralAssetAddress)
		assert.Equal(t, "100.0000000", result[1].RequiredAmount)
		assert.Equal(t, "300.0000000", result[1].PresentAmount)
	})

	t.Run("fail, call Heart Contract GetStableCreditCount", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(0), errors.New("error here"))

		result, err := testHelper.Client.CollateralHealthCheck(&CollateralHealthCheckInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call Heart Contract GetStableCreditCount with stableCreditSize 0", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(0), nil)

		result, err := testHelper.Client.CollateralHealthCheck(&CollateralHealthCheckInput{})

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

		result, err := testHelper.Client.CollateralHealthCheck(&CollateralHealthCheckInput{})

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

		result, err := testHelper.Client.CollateralHealthCheck(&CollateralHealthCheckInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call Drs Contract CollateralHealthCheck 1 loop", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		requiredAmount, _ := utils.StringToAmount("")
		presentAmount, _ := utils.StringToAmount("")
		collateralAssetCode := utils.StringToByte32("")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"
		toAbiInput := &CollateralHealthCheckAbiInput{
			AssetCode: stableCreditAssetCode,
		}

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(1), nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"), nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"))).
			Return(&stableCreditAssetCode, &stableCreditId, nil)

		testHelper.MockDRSContract.EXPECT().
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(common.Address{}, collateralAssetCode, requiredAmount, presentAmount, errors.New("error here"))

		result, err := testHelper.Client.CollateralHealthCheck(&CollateralHealthCheckInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call Drs Contract CollateralHealthCheck 2 loop", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		requiredAmount, _ := utils.StringToAmount("100")
		presentAmount, _ := utils.StringToAmount("200")
		collateralAssetCode := utils.StringToByte32("VELO")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"
		toAbiInput := &CollateralHealthCheckAbiInput{
			AssetCode: stableCreditAssetCode,
		}

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(2), nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"), nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"))).
			Return(&stableCreditAssetCode, &stableCreditId, nil)
		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"))).
			Return(&stableCreditAssetCode, &stableCreditId, nil)

		testHelper.MockHeartContract.EXPECT().
			GetNextStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{}), gomock.AssignableToTypeOf(stableCreditId)).
			Return(common.HexToAddress("0x8F384Dd497A06941531ED277Cc0Cd34e7d594323"), nil)

		testHelper.MockDRSContract.EXPECT().
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(common.Address{}, collateralAssetCode, requiredAmount, presentAmount, nil)
		testHelper.MockDRSContract.EXPECT().
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(common.Address{}, collateralAssetCode, requiredAmount, presentAmount, errors.New("error here"))

		result, err := testHelper.Client.CollateralHealthCheck(&CollateralHealthCheckInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call Drs Contract CollateralHealthCheck 2 loop but got error from call txHelper StableCreditAssetCode", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		requiredAmount, _ := utils.StringToAmount("100")
		presentAmount, _ := utils.StringToAmount("200")
		collateralAssetCode := utils.StringToByte32("VELO")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"
		toAbiInput := &CollateralHealthCheckAbiInput{
			AssetCode: stableCreditAssetCode,
		}

		testHelper.MockHeartContract.EXPECT().
			GetStableCreditCount(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(uint8(2), nil)

		testHelper.MockHeartContract.EXPECT().
			GetRecentStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{})).
			Return(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"), nil)

		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"))).
			Return(&stableCreditAssetCode, &stableCreditId, nil)
		testHelper.MockTxHelper.EXPECT().
			StableCreditAssetCode(gomock.AssignableToTypeOf(common.HexToAddress("0x2Fb9287ba799297EB918aD607B5F8D3108a026f0"))).
			Return(nil, nil, errors.New("error here"))

		testHelper.MockHeartContract.EXPECT().
			GetNextStableCredit(gomock.AssignableToTypeOf(&bind.CallOpts{}), gomock.AssignableToTypeOf(stableCreditId)).
			Return(common.HexToAddress("0x8F384Dd497A06941531ED277Cc0Cd34e7d594323"), nil)

		testHelper.MockDRSContract.EXPECT().
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(common.Address{}, collateralAssetCode, requiredAmount, presentAmount, nil)

		result, err := testHelper.Client.CollateralHealthCheck(&CollateralHealthCheckInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fail, call Drs Contract CollateralHealthCheck 2 loop but got error from call Heart Contract GetNextStableCredit", func(t *testing.T) {
		testHelper := testHelperWithMock(t)
		defer testHelper.MockController.Finish()

		requiredAmount, _ := utils.StringToAmount("100")
		presentAmount, _ := utils.StringToAmount("200")
		collateralAssetCode := utils.StringToByte32("VELO")
		stableCreditId := utils.StringToByte32("VELO_vUSD_0x2Fb9287ba799297EB918aD607B5F8D3108a026f0")
		stableCreditAssetCode := "vUSD"
		toAbiInput := &CollateralHealthCheckAbiInput{
			AssetCode: stableCreditAssetCode,
		}

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
			CollateralHealthCheck(gomock.AssignableToTypeOf(&bind.CallOpts{}), toAbiInput.AssetCode).
			Return(common.Address{}, collateralAssetCode, requiredAmount, presentAmount, nil)

		result, err := testHelper.Client.CollateralHealthCheck(&CollateralHealthCheckInput{})

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
