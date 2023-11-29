package node

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

func repayExcessSD(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Check and assign the EC status
	err = cliutils.CheckClientStatus(staderClient)
	if err != nil {
		return err
	}

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)
	if err != nil {
		return err
	}

	// Get stake mount
	amountInString := c.String("amount")

	amount, err := strconv.ParseFloat(amountInString, 64)
	if err != nil {
		return err
	}

	amountWei := eth.EthToWei(amount)

	canRepayExcessSdResponse, err := staderClient.CanRepayExcessSD(amountWei)
	if err != nil {
		return err
	}

	sdStatus := canRepayExcessSdResponse.SdStatusResponse

	// Less than 200 %
	if sdStatus.SdCollateralCurrentAmount.Cmp(sdStatus.SdMaxCollateralAmount) < 0 {
		fmt.Printf("Not enough SD collateral to repay utilized SD %s \n", sdStatus.PoolAvailableSDBalance.String())
		return nil
	}

	// Do not had position
	if sdStatus.SdUtilizerLatestBalance.Cmp(big.NewInt(0)) <= 0 {
		fmt.Printf("You don't have an existing utilization position. To withdraw excess SD to your wallet execute the following command stader-cli node withdraw-sd\n")
		return nil
	}

	amountToRepayExcess := new(big.Int).Sub(sdStatus.SdUtilizerLatestBalance, sdStatus.SdMaxCollateralAmount)

	err = gas.AssignMaxFeeAndLimit(canRepayExcessSdResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintln(
		"Are you sure you want to repay SD?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.NodeRepaySd(amountToRepayExcess)
	if err != nil {
		return err
	}

	cliutils.PrintTransactionHash(staderClient, res.TxHash)

	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	return nil
}
