package stader

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-node/shared/types/api"
)

// Get node status
func (c *Client) NodeStatus() (api.NodeStatusResponse, error) {
	responseBytes, err := c.callAPI("node status")
	if err != nil {
		return api.NodeStatusResponse{}, fmt.Errorf("Could not get node status: %w", err)
	}
	var response api.NodeStatusResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStatusResponse{}, fmt.Errorf("Could not decode node status response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStatusResponse{}, fmt.Errorf("Could not get node status: %s", response.Error)
	}
	if response.AccountBalances.ETH == nil {
		response.AccountBalances.ETH = big.NewInt(0)
	}
	if response.AccountBalances.Sd == nil {
		response.AccountBalances.Sd = big.NewInt(0)
	}

	return response, nil
}

// Check whether the node can be registered
func (c *Client) CanRegisterNode(operatorName string, operatorRewardAddress common.Address, socializeMev bool) (api.CanRegisterNodeResponse, error) {
	responseBytes, err := c.callAPI("node can-register", operatorName, operatorRewardAddress.Hex(), strconv.FormatBool(socializeMev))
	if err != nil {
		return api.CanRegisterNodeResponse{}, fmt.Errorf("Could not get can register node status: %w", err)
	}
	var response api.CanRegisterNodeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanRegisterNodeResponse{}, fmt.Errorf("Could not decode can register node response: %w", err)
	}
	if response.Error != "" {
		return api.CanRegisterNodeResponse{}, fmt.Errorf("Could not get can register node status: %s", response.Error)
	}
	return response, nil
}

// Register the node
func (c *Client) RegisterNode(operatorName string, operatorRewardAddress common.Address, socializeMev bool) (api.RegisterNodeResponse, error) {
	responseBytes, err := c.callAPI("node register", operatorName, operatorRewardAddress.Hex(), strconv.FormatBool(socializeMev))
	if err != nil {
		return api.RegisterNodeResponse{}, fmt.Errorf("Could not register node: %w", err)
	}
	var response api.RegisterNodeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.RegisterNodeResponse{}, fmt.Errorf("Could not decode register node response: %w", err)
	}
	if response.Error != "" {
		return api.RegisterNodeResponse{}, fmt.Errorf("Could not register node: %s", response.Error)
	}
	return response, nil
}

// Checks if the node's withdrawal address can be set
func (c *Client) CanSetNodeWithdrawalAddress(withdrawalAddress common.Address, confirm bool) (api.CanSetNodeWithdrawalAddressResponse, error) {
	responseBytes, err := c.callAPI("node can-set-withdrawal-address", withdrawalAddress.Hex(), strconv.FormatBool(confirm))
	if err != nil {
		return api.CanSetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not set node withdrawal address: %w", err)
	}
	var response api.CanSetNodeWithdrawalAddressResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanSetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not decode can set node withdrawal address response: %w", err)
	}
	if response.Error != "" {
		return api.CanSetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not set node withdrawal address: %s", response.Error)
	}
	return response, nil
}

// Set the node's withdrawal address
func (c *Client) SetNodeWithdrawalAddress(withdrawalAddress common.Address, confirm bool) (api.SetNodeWithdrawalAddressResponse, error) {
	responseBytes, err := c.callAPI("node set-withdrawal-address", withdrawalAddress.Hex(), strconv.FormatBool(confirm))
	if err != nil {
		return api.SetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not set node withdrawal address: %w", err)
	}
	var response api.SetNodeWithdrawalAddressResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not decode set node withdrawal address response: %w", err)
	}
	if response.Error != "" {
		return api.SetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not set node withdrawal address: %s", response.Error)
	}
	return response, nil
}

// Checks if the node's withdrawal address can be confirmed
func (c *Client) CanConfirmNodeWithdrawalAddress() (api.CanSetNodeWithdrawalAddressResponse, error) {
	responseBytes, err := c.callAPI("node can-confirm-withdrawal-address")
	if err != nil {
		return api.CanSetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not get can confirm node withdrawal address: %w", err)
	}
	var response api.CanSetNodeWithdrawalAddressResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanSetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not decode can confirm node withdrawal address response: %w", err)
	}
	if response.Error != "" {
		return api.CanSetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not get can confirm node withdrawal address: %s", response.Error)
	}
	return response, nil
}

// Confirm the node's withdrawal address
func (c *Client) ConfirmNodeWithdrawalAddress() (api.SetNodeWithdrawalAddressResponse, error) {
	responseBytes, err := c.callAPI("node confirm-withdrawal-address")
	if err != nil {
		return api.SetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not confirm node withdrawal address: %w", err)
	}
	var response api.SetNodeWithdrawalAddressResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not decode confirm node withdrawal address response: %w", err)
	}
	if response.Error != "" {
		return api.SetNodeWithdrawalAddressResponse{}, fmt.Errorf("Could not confirm node withdrawal address: %s", response.Error)
	}
	return response, nil
}

// Checks if the node's timezone location can be set
func (c *Client) CanSetNodeTimezone(timezoneLocation string) (api.CanSetNodeTimezoneResponse, error) {
	responseBytes, err := c.callAPI("node can-set-timezone", timezoneLocation)
	if err != nil {
		return api.CanSetNodeTimezoneResponse{}, fmt.Errorf("Could not get can set node timezone: %w", err)
	}
	var response api.CanSetNodeTimezoneResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanSetNodeTimezoneResponse{}, fmt.Errorf("Could not decode can set node timezone response: %w", err)
	}
	if response.Error != "" {
		return api.CanSetNodeTimezoneResponse{}, fmt.Errorf("Could not get can set node timezone: %s", response.Error)
	}
	return response, nil
}

// Set the node's timezone location
func (c *Client) SetNodeTimezone(timezoneLocation string) (api.SetNodeTimezoneResponse, error) {
	responseBytes, err := c.callAPI("node set-timezone", timezoneLocation)
	if err != nil {
		return api.SetNodeTimezoneResponse{}, fmt.Errorf("Could not set node timezone: %w", err)
	}
	var response api.SetNodeTimezoneResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SetNodeTimezoneResponse{}, fmt.Errorf("Could not decode set node timezone response: %w", err)
	}
	if response.Error != "" {
		return api.SetNodeTimezoneResponse{}, fmt.Errorf("Could not set node timezone: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can swap RPL tokens
func (c *Client) CanNodeSwapRpl(amountWei *big.Int) (api.CanNodeSwapRplResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-swap-rpl %s", amountWei.String()))
	if err != nil {
		return api.CanNodeSwapRplResponse{}, fmt.Errorf("Could not get can node swap RPL status: %w", err)
	}
	var response api.CanNodeSwapRplResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeSwapRplResponse{}, fmt.Errorf("Could not decode can node swap RPL response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeSwapRplResponse{}, fmt.Errorf("Could not get can node swap RPL status: %s", response.Error)
	}
	return response, nil
}

// Get the gas estimate for approving legacy RPL interaction
func (c *Client) NodeSwapRplApprovalGas(amountWei *big.Int) (api.NodeSwapRplApproveGasResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node get-swap-rpl-approval-gas %s", amountWei.String()))
	if err != nil {
		return api.NodeSwapRplApproveGasResponse{}, fmt.Errorf("Could not get old RPL approval gas: %w", err)
	}
	var response api.NodeSwapRplApproveGasResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSwapRplApproveGasResponse{}, fmt.Errorf("Could not decode node swap RPL approve gas response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSwapRplApproveGasResponse{}, fmt.Errorf("Could not get old RPL approval gas: %s", response.Error)
	}
	return response, nil
}

// Approves old RPL for a token swap
func (c *Client) NodeSwapRplApprove(amountWei *big.Int) (api.NodeSwapRplApproveResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node swap-rpl-approve-rpl %s", amountWei.String()))
	if err != nil {
		return api.NodeSwapRplApproveResponse{}, fmt.Errorf("Could not approve old RPL: %w", err)
	}
	var response api.NodeSwapRplApproveResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSwapRplApproveResponse{}, fmt.Errorf("Could not decode node swap RPL approve response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSwapRplApproveResponse{}, fmt.Errorf("Could not approve old RPL tokens for swapping: %s", response.Error)
	}
	return response, nil
}

// Swap node's old RPL tokens for new RPL tokens, waiting for the approval to be included in a block first
func (c *Client) NodeWaitAndSwapRpl(amountWei *big.Int, approvalTxHash common.Hash) (api.NodeSwapRplSwapResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node wait-and-swap-rpl %s %s", amountWei.String(), approvalTxHash.String()))
	if err != nil {
		return api.NodeSwapRplSwapResponse{}, fmt.Errorf("Could not swap node's RPL tokens: %w", err)
	}
	var response api.NodeSwapRplSwapResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSwapRplSwapResponse{}, fmt.Errorf("Could not decode node swap RPL tokens response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSwapRplSwapResponse{}, fmt.Errorf("Could not swap node's RPL tokens: %s", response.Error)
	}
	return response, nil
}

// Swap node's old RPL tokens for new RPL tokens
func (c *Client) NodeSwapRpl(amountWei *big.Int) (api.NodeSwapRplSwapResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node swap-rpl %s", amountWei.String()))
	if err != nil {
		return api.NodeSwapRplSwapResponse{}, fmt.Errorf("Could not swap node's RPL tokens: %w", err)
	}
	var response api.NodeSwapRplSwapResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSwapRplSwapResponse{}, fmt.Errorf("Could not decode node swap RPL tokens response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSwapRplSwapResponse{}, fmt.Errorf("Could not swap node's RPL tokens: %s", response.Error)
	}
	return response, nil
}

// Get a node's legacy RPL allowance for swapping on the new RPL contract
func (c *Client) GetNodeSwapRplAllowance() (api.NodeSwapRplAllowanceResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node swap-rpl-allowance"))
	if err != nil {
		return api.NodeSwapRplAllowanceResponse{}, fmt.Errorf("Could not get node swap RPL allowance: %w", err)
	}
	var response api.NodeSwapRplAllowanceResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSwapRplAllowanceResponse{}, fmt.Errorf("Could not decode node swap RPL allowance response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSwapRplAllowanceResponse{}, fmt.Errorf("Could not get node swap RPL allowance: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can stake RPL
func (c *Client) CanNodeDepositSd(amountWei *big.Int) (api.CanNodeStakeRplResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-node-deposit-sd %s", amountWei.String()))
	if err != nil {
		return api.CanNodeStakeRplResponse{}, fmt.Errorf("Could not get can node stake RPL status: %w", err)
	}
	var response api.CanNodeStakeRplResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeStakeRplResponse{}, fmt.Errorf("Could not decode can node stake RPL response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeStakeRplResponse{}, fmt.Errorf("Could not get can node stake RPL status: %s", response.Error)
	}
	return response, nil
}

// Get the gas estimate for approving new RPL interaction
func (c *Client) NodeDepositSdApprovalGas(amountWei *big.Int) (api.NodeStakeRplApproveGasResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node get-deposit-sd-approval-gas %s", amountWei.String()))
	if err != nil {
		return api.NodeStakeRplApproveGasResponse{}, fmt.Errorf("Could not get new RPL approval gas: %w", err)
	}
	var response api.NodeStakeRplApproveGasResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStakeRplApproveGasResponse{}, fmt.Errorf("Could not decode node stake RPL approve gas response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStakeRplApproveGasResponse{}, fmt.Errorf("Could not get new RPL approval gas: %s", response.Error)
	}
	return response, nil
}

// Approve RPL for staking against the node
func (c *Client) NodeDepositSdApprove(amountWei *big.Int) (api.NodeStakeRplApproveResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit-sd-approve-sd %s", amountWei.String()))
	if err != nil {
		return api.NodeStakeRplApproveResponse{}, fmt.Errorf("Could not approve RPL for staking: %w", err)
	}
	var response api.NodeStakeRplApproveResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStakeRplApproveResponse{}, fmt.Errorf("Could not decode stake node RPL approve response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStakeRplApproveResponse{}, fmt.Errorf("Could not approve RPL for staking: %s", response.Error)
	}
	return response, nil
}

// Stake RPL against the node waiting for approvalTxHash to be included in a block first
func (c *Client) NodeWaitAndStakeRpl(amountWei *big.Int, approvalTxHash common.Hash) (api.NodeStakeRplStakeResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node wait-and-stake-rpl %s %s", amountWei.String(), approvalTxHash.String()))
	if err != nil {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("Could not stake node RPL: %w", err)
	}
	var response api.NodeStakeRplStakeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("Could not decode stake node RPL response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("Could not stake node RPL: %s", response.Error)
	}
	return response, nil
}

// Stake RPL against the node
func (c *Client) NodeDepositSd(amountWei *big.Int) (api.NodeStakeRplStakeResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit-sd %s", amountWei.String()))
	if err != nil {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("Could not stake node RPL: %w", err)
	}
	var response api.NodeStakeRplStakeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("Could not decode stake node RPL response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStakeRplStakeResponse{}, fmt.Errorf("Could not stake node RPL: %s", response.Error)
	}
	return response, nil
}

// Get a node's RPL allowance for the staking contract
func (c *Client) GetNodeDepositSdAllowance() (api.NodeStakeRplAllowanceResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit-sd-allowance"))
	if err != nil {
		return api.NodeStakeRplAllowanceResponse{}, fmt.Errorf("Could not get node stake RPL allowance: %w", err)
	}
	var response api.NodeStakeRplAllowanceResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStakeRplAllowanceResponse{}, fmt.Errorf("Could not decode node stake RPL allowance response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStakeRplAllowanceResponse{}, fmt.Errorf("Could not get node stake RPL allowance: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can withdraw RPL
func (c *Client) CanNodeWithdrawRpl(amountWei *big.Int) (api.CanNodeWithdrawRplResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-withdraw-rpl %s", amountWei.String()))
	if err != nil {
		return api.CanNodeWithdrawRplResponse{}, fmt.Errorf("Could not get can node withdraw RPL status: %w", err)
	}
	var response api.CanNodeWithdrawRplResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeWithdrawRplResponse{}, fmt.Errorf("Could not decode can node withdraw RPL response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeWithdrawRplResponse{}, fmt.Errorf("Could not get can node withdraw RPL status: %s", response.Error)
	}
	return response, nil
}

// Withdraw RPL staked against the node
func (c *Client) NodeWithdrawRpl(amountWei *big.Int) (api.NodeWithdrawRplResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node withdraw-rpl %s", amountWei.String()))
	if err != nil {
		return api.NodeWithdrawRplResponse{}, fmt.Errorf("Could not withdraw node RPL: %w", err)
	}
	var response api.NodeWithdrawRplResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeWithdrawRplResponse{}, fmt.Errorf("Could not decode withdraw node RPL response: %w", err)
	}
	if response.Error != "" {
		return api.NodeWithdrawRplResponse{}, fmt.Errorf("Could not withdraw node RPL: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can make a deposit
func (c *Client) CanNodeDeposit(amountWei *big.Int, salt *big.Int, numValidators *big.Int, submit bool) (api.CanNodeDepositResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-deposit %s %s %d %t", amountWei.String(), salt.String(), numValidators, submit))
	if err != nil {
		return api.CanNodeDepositResponse{}, fmt.Errorf("Could not get can node deposit status: %w", err)
	}
	var response api.CanNodeDepositResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeDepositResponse{}, fmt.Errorf("Could not decode can node deposit response: %w", err)
	}
	fmt.Printf("response in CanNodeDeposit is %v\n", response)
	if response.Error != "" {
		return api.CanNodeDepositResponse{}, fmt.Errorf("Could not get can node deposit status: %s", response.Error)
	}
	return response, nil
}

// Make a node deposit
func (c *Client) NodeDeposit(amountWei *big.Int, salt *big.Int, numValidators *big.Int, submit bool) (api.NodeDepositResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node deposit %s %s %d %t", amountWei.String(), salt.String(), numValidators, submit))
	if err != nil {
		return api.NodeDepositResponse{}, fmt.Errorf("Could not make node deposit as er: %w", err)
	}
	var response api.NodeDepositResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeDepositResponse{}, fmt.Errorf("Could not decode node deposit response: %w", err)
	}
	if response.Error != "" {
		return api.NodeDepositResponse{}, fmt.Errorf("Could not make node deposit: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can send tokens
func (c *Client) CanNodeSend(amountWei *big.Int, token string) (api.CanNodeSendResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-send %s %s", amountWei.String(), token))
	if err != nil {
		return api.CanNodeSendResponse{}, fmt.Errorf("Could not get can node send status: %w", err)
	}
	var response api.CanNodeSendResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeSendResponse{}, fmt.Errorf("Could not decode can node send response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeSendResponse{}, fmt.Errorf("Could not get can node send status: %s", response.Error)
	}
	return response, nil
}

// Send tokens from the node to an address
func (c *Client) NodeSend(amountWei *big.Int, token string, toAddress common.Address) (api.NodeSendResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node send %s %s %s", amountWei.String(), token, toAddress.Hex()))
	if err != nil {
		return api.NodeSendResponse{}, fmt.Errorf("Could not send tokens from node: %w", err)
	}
	var response api.NodeSendResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSendResponse{}, fmt.Errorf("Could not decode node send response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSendResponse{}, fmt.Errorf("Could not send tokens from node: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can burn tokens
func (c *Client) CanNodeBurn(amountWei *big.Int, token string) (api.CanNodeBurnResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-burn %s %s", amountWei.String(), token))
	if err != nil {
		return api.CanNodeBurnResponse{}, fmt.Errorf("Could not get can node burn status: %w", err)
	}
	var response api.CanNodeBurnResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeBurnResponse{}, fmt.Errorf("Could not decode can node burn response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeBurnResponse{}, fmt.Errorf("Could not get can node burn status: %s", response.Error)
	}
	return response, nil
}

// Burn tokens owned by the node for ETH
func (c *Client) NodeBurn(amountWei *big.Int, token string) (api.NodeBurnResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node burn %s %s", amountWei.String(), token))
	if err != nil {
		return api.NodeBurnResponse{}, fmt.Errorf("Could not burn tokens owned by node: %w", err)
	}
	var response api.NodeBurnResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeBurnResponse{}, fmt.Errorf("Could not decode node burn response: %w", err)
	}
	if response.Error != "" {
		return api.NodeBurnResponse{}, fmt.Errorf("Could not burn tokens owned by node: %s", response.Error)
	}
	return response, nil
}

// Get node sync progress
func (c *Client) NodeSync() (api.NodeSyncProgressResponse, error) {
	responseBytes, err := c.callAPI("node sync")
	if err != nil {
		return api.NodeSyncProgressResponse{}, fmt.Errorf("Could not get node sync: %w", err)
	}
	var response api.NodeSyncProgressResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSyncProgressResponse{}, fmt.Errorf("Could not decode node sync response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSyncProgressResponse{}, fmt.Errorf("Could not get node sync: %s", response.Error)
	}
	return response, nil
}

// Check whether the node has RPL rewards available to claim
func (c *Client) CanNodeClaimRpl() (api.CanNodeClaimRplResponse, error) {
	responseBytes, err := c.callAPI("node can-claim-rpl-rewards")
	if err != nil {
		return api.CanNodeClaimRplResponse{}, fmt.Errorf("Could not get can node claim rpl rewards status: %w", err)
	}
	var response api.CanNodeClaimRplResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeClaimRplResponse{}, fmt.Errorf("Could not decode can node claim rpl rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeClaimRplResponse{}, fmt.Errorf("Could not get can node claim rpl rewards status: %s", response.Error)
	}
	return response, nil
}

// Claim available RPL rewards
func (c *Client) NodeClaimRpl() (api.NodeClaimRplResponse, error) {
	responseBytes, err := c.callAPI("node claim-rpl-rewards")
	if err != nil {
		return api.NodeClaimRplResponse{}, fmt.Errorf("Could not claim rpl rewards: %w", err)
	}
	var response api.NodeClaimRplResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeClaimRplResponse{}, fmt.Errorf("Could not decode node claim rpl rewards response: %w", err)
	}
	if response.Error != "" {
		return api.NodeClaimRplResponse{}, fmt.Errorf("Could not claim rpl rewards: %s", response.Error)
	}
	return response, nil
}

// Get node RPL rewards status
func (c *Client) NodeRewards() (api.NodeRewardsResponse, error) {
	responseBytes, err := c.callAPI("node rewards")
	if err != nil {
		return api.NodeRewardsResponse{}, fmt.Errorf("Could not get node rewards: %w", err)
	}
	var response api.NodeRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeRewardsResponse{}, fmt.Errorf("Could not decode node rewards response: %w", err)
	}
	if response.Error != "" {
		return api.NodeRewardsResponse{}, fmt.Errorf("Could not get node rewards: %s", response.Error)
	}
	return response, nil
}

// Get the deposit contract info for Stader and the Beacon Client
func (c *Client) DepositContractInfo() (api.ContractsInfoResponse, error) {
	responseBytes, err := c.callAPI("node deposit-contract-info")
	if err != nil {
		return api.ContractsInfoResponse{}, fmt.Errorf("Could not get deposit contract info: %w", err)
	}
	var response api.ContractsInfoResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ContractsInfoResponse{}, fmt.Errorf("Could not decode deposit contract info response: %w", err)
	}
	if response.Error != "" {
		return api.ContractsInfoResponse{}, fmt.Errorf("Could not get deposit contract info: %s", response.Error)
	}
	return response, nil
}

// Estimate the gas required to set a voting snapshot delegate
func (c *Client) EstimateSetSnapshotDelegateGas(address common.Address) (api.EstimateSetSnapshotDelegateGasResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node estimate-set-snapshot-delegate-gas %s", address.Hex()))
	if err != nil {
		return api.EstimateSetSnapshotDelegateGasResponse{}, fmt.Errorf("Could not get estimate-set-snapshot-delegate-gas response: %w", err)
	}
	var response api.EstimateSetSnapshotDelegateGasResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.EstimateSetSnapshotDelegateGasResponse{}, fmt.Errorf("Could not decode estimate-set-snapshot-delegate-gas response: %w", err)
	}
	if response.Error != "" {
		return api.EstimateSetSnapshotDelegateGasResponse{}, fmt.Errorf("Could not get estimate-set-snapshot-delegate-gas response: %s", response.Error)
	}
	return response, nil
}

// Set a voting snapshot delegate for the node
func (c *Client) SetSnapshotDelegate(address common.Address) (api.SetSnapshotDelegateResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node set-snapshot-delegate %s", address.Hex()))
	if err != nil {
		return api.SetSnapshotDelegateResponse{}, fmt.Errorf("Could not get set-snapshot-delegate response: %w", err)
	}
	var response api.SetSnapshotDelegateResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SetSnapshotDelegateResponse{}, fmt.Errorf("Could not decode set-snapshot-delegate response: %w", err)
	}
	if response.Error != "" {
		return api.SetSnapshotDelegateResponse{}, fmt.Errorf("Could not get set-snapshot-delegate response: %s", response.Error)
	}
	return response, nil
}

// Estimate the gas required to clear the node's voting snapshot delegate
func (c *Client) EstimateClearSnapshotDelegateGas() (api.EstimateClearSnapshotDelegateGasResponse, error) {
	responseBytes, err := c.callAPI("node estimate-clear-snapshot-delegate-gas")
	if err != nil {
		return api.EstimateClearSnapshotDelegateGasResponse{}, fmt.Errorf("Could not get estimate-clear-snapshot-delegate-gas response: %w", err)
	}
	var response api.EstimateClearSnapshotDelegateGasResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.EstimateClearSnapshotDelegateGasResponse{}, fmt.Errorf("Could not decode estimate-clear-snapshot-delegate-gas response: %w", err)
	}
	if response.Error != "" {
		return api.EstimateClearSnapshotDelegateGasResponse{}, fmt.Errorf("Could not get estimate-clear-snapshot-delegate-gas response: %s", response.Error)
	}
	return response, nil
}

// Clear the node's voting snapshot delegate
func (c *Client) ClearSnapshotDelegate() (api.ClearSnapshotDelegateResponse, error) {
	responseBytes, err := c.callAPI("node clear-snapshot-delegate")
	if err != nil {
		return api.ClearSnapshotDelegateResponse{}, fmt.Errorf("Could not get clear-snapshot-delegate response: %w", err)
	}
	var response api.ClearSnapshotDelegateResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ClearSnapshotDelegateResponse{}, fmt.Errorf("Could not decode clear-snapshot-delegate response: %w", err)
	}
	if response.Error != "" {
		return api.ClearSnapshotDelegateResponse{}, fmt.Errorf("Could not get clear-snapshot-delegate response: %s", response.Error)
	}
	return response, nil
}

// Get the initialization status of the fee distributor contract
func (c *Client) IsFeeDistributorInitialized() (api.NodeIsFeeDistributorInitializedResponse, error) {
	responseBytes, err := c.callAPI("node is-fee-distributor-initialized")
	if err != nil {
		return api.NodeIsFeeDistributorInitializedResponse{}, fmt.Errorf("Could not get fee distributor initialization status: %w", err)
	}
	var response api.NodeIsFeeDistributorInitializedResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeIsFeeDistributorInitializedResponse{}, fmt.Errorf("Could not decode fee distributor initialization status response: %w", err)
	}
	if response.Error != "" {
		return api.NodeIsFeeDistributorInitializedResponse{}, fmt.Errorf("Could not get fee distributor initialization status: %s", response.Error)
	}
	return response, nil
}

// Get the gas cost for initializing the fee distributor contract
func (c *Client) GetInitializeFeeDistributorGas() (api.NodeInitializeFeeDistributorGasResponse, error) {
	responseBytes, err := c.callAPI("node get-initialize-fee-distributor-gas")
	if err != nil {
		return api.NodeInitializeFeeDistributorGasResponse{}, fmt.Errorf("Could not get initialize fee distributor gas: %w", err)
	}
	var response api.NodeInitializeFeeDistributorGasResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeInitializeFeeDistributorGasResponse{}, fmt.Errorf("Could not decode initialize fee distributor gas response: %w", err)
	}
	if response.Error != "" {
		return api.NodeInitializeFeeDistributorGasResponse{}, fmt.Errorf("Could not get initialize fee distributor gas: %s", response.Error)
	}
	return response, nil
}

// Initialize the fee distributor contract
func (c *Client) InitializeFeeDistributor() (api.NodeInitializeFeeDistributorResponse, error) {
	responseBytes, err := c.callAPI("node initialize-fee-distributor")
	if err != nil {
		return api.NodeInitializeFeeDistributorResponse{}, fmt.Errorf("Could not initialize fee distributor: %w", err)
	}
	var response api.NodeInitializeFeeDistributorResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeInitializeFeeDistributorResponse{}, fmt.Errorf("Could not decode initialize fee distributor response: %w", err)
	}
	if response.Error != "" {
		return api.NodeInitializeFeeDistributorResponse{}, fmt.Errorf("Could not initialize fee distributor: %s", response.Error)
	}
	return response, nil
}

// Check if distributing ETH from the node's fee distributor is possible
func (c *Client) CanDistribute() (api.NodeCanDistributeResponse, error) {
	responseBytes, err := c.callAPI("node can-distribute")
	if err != nil {
		return api.NodeCanDistributeResponse{}, fmt.Errorf("Could not get can distribute: %w", err)
	}
	var response api.NodeCanDistributeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeCanDistributeResponse{}, fmt.Errorf("Could not decode can distribute response: %w", err)
	}
	if response.Error != "" {
		return api.NodeCanDistributeResponse{}, fmt.Errorf("Could not get can distribute: %s", response.Error)
	}
	return response, nil
}

// Distribute ETH from the node's fee distributor
func (c *Client) Distribute() (api.NodeDistributeResponse, error) {
	responseBytes, err := c.callAPI("node distribute")
	if err != nil {
		return api.NodeDistributeResponse{}, fmt.Errorf("Could not distribute ETH: %w", err)
	}
	var response api.NodeDistributeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeDistributeResponse{}, fmt.Errorf("Could not decode distribute response: %w", err)
	}
	if response.Error != "" {
		return api.NodeDistributeResponse{}, fmt.Errorf("Could not distribute ETH: %s", response.Error)
	}
	return response, nil
}

// Get info about your eligible rewards periods, including balances and Merkle proofs
func (c *Client) GetRewardsInfo() (api.NodeGetRewardsInfoResponse, error) {
	responseBytes, err := c.callAPI("node get-rewards-info")
	if err != nil {
		return api.NodeGetRewardsInfoResponse{}, fmt.Errorf("Could not get rewards info: %w", err)
	}
	var response api.NodeGetRewardsInfoResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeGetRewardsInfoResponse{}, fmt.Errorf("Could not decode get rewards info response: %w", err)
	}
	if response.Error != "" {
		return api.NodeGetRewardsInfoResponse{}, fmt.Errorf("Could not get rewards info: %s", response.Error)
	}
	return response, nil
}

// Check if the rewards for the given intervals can be claimed
func (c *Client) CanNodeClaimRewards(indices []uint64) (api.CanNodeClaimRewardsResponse, error) {
	indexStrings := []string{}
	for _, index := range indices {
		indexStrings = append(indexStrings, fmt.Sprint(index))
	}
	responseBytes, err := c.callAPI("node can-claim-rewards", strings.Join(indexStrings, ","))
	if err != nil {
		return api.CanNodeClaimRewardsResponse{}, fmt.Errorf("Could not check if can claim rewards: %w", err)
	}
	var response api.CanNodeClaimRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeClaimRewardsResponse{}, fmt.Errorf("Could not decode can claim rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeClaimRewardsResponse{}, fmt.Errorf("Could not check if can claim rewards: %s", response.Error)
	}
	return response, nil
}

// Claim rewards for the given reward intervals
func (c *Client) NodeClaimRewards(indices []uint64) (api.NodeClaimRewardsResponse, error) {
	indexStrings := []string{}
	for _, index := range indices {
		indexStrings = append(indexStrings, fmt.Sprint(index))
	}
	responseBytes, err := c.callAPI("node claim-rewards", strings.Join(indexStrings, ","))
	if err != nil {
		return api.NodeClaimRewardsResponse{}, fmt.Errorf("Could not claim rewards: %w", err)
	}
	var response api.NodeClaimRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeClaimRewardsResponse{}, fmt.Errorf("Could not decode claim rewards response: %w", err)
	}
	if response.Error != "" {
		return api.NodeClaimRewardsResponse{}, fmt.Errorf("Could not claim rewards: %s", response.Error)
	}
	return response, nil
}

// Check if the rewards for the given intervals can be claimed, and RPL restaked automatically
func (c *Client) CanNodeClaimAndStakeRewards(indices []uint64, stakeAmountWei *big.Int) (api.CanNodeClaimAndStakeRewardsResponse, error) {
	indexStrings := []string{}
	for _, index := range indices {
		indexStrings = append(indexStrings, fmt.Sprint(index))
	}
	responseBytes, err := c.callAPI("node can-claim-and-stake-rewards", strings.Join(indexStrings, ","), stakeAmountWei.String())
	if err != nil {
		return api.CanNodeClaimAndStakeRewardsResponse{}, fmt.Errorf("Could not check if can claim and stake rewards: %w", err)
	}
	var response api.CanNodeClaimAndStakeRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeClaimAndStakeRewardsResponse{}, fmt.Errorf("Could not decode can claim and stake rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeClaimAndStakeRewardsResponse{}, fmt.Errorf("Could not check if can claim and stake rewards: %s", response.Error)
	}
	return response, nil
}

// Claim rewards for the given reward intervals and restake RPL automatically
func (c *Client) NodeClaimAndStakeRewards(indices []uint64, stakeAmountWei *big.Int) (api.NodeClaimAndStakeRewardsResponse, error) {
	indexStrings := []string{}
	for _, index := range indices {
		indexStrings = append(indexStrings, fmt.Sprint(index))
	}
	responseBytes, err := c.callAPI("node claim-and-stake-rewards", strings.Join(indexStrings, ","), stakeAmountWei.String())
	if err != nil {
		return api.NodeClaimAndStakeRewardsResponse{}, fmt.Errorf("Could not claim and stake rewards: %w", err)
	}
	var response api.NodeClaimAndStakeRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeClaimAndStakeRewardsResponse{}, fmt.Errorf("Could not decode claim and stake rewards response: %w", err)
	}
	if response.Error != "" {
		return api.NodeClaimAndStakeRewardsResponse{}, fmt.Errorf("Could not claim and stake rewards: %s", response.Error)
	}
	return response, nil
}

// Check whether or not the node is opted into the Smoothing Pool
func (c *Client) NodeGetSmoothingPoolRegistrationStatus() (api.GetSmoothingPoolRegistrationStatusResponse, error) {
	responseBytes, err := c.callAPI("node get-smoothing-pool-registration-status")
	if err != nil {
		return api.GetSmoothingPoolRegistrationStatusResponse{}, fmt.Errorf("Could not get smoothing pool registration status: %w", err)
	}
	var response api.GetSmoothingPoolRegistrationStatusResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.GetSmoothingPoolRegistrationStatusResponse{}, fmt.Errorf("Could not decode smoothing pool registration status response: %w", err)
	}
	if response.Error != "" {
		return api.GetSmoothingPoolRegistrationStatusResponse{}, fmt.Errorf("Could not get smoothing pool registration status: %s", response.Error)
	}
	return response, nil
}

// Check if the node's Smoothing Pool status can be changed
func (c *Client) CanNodeSetSmoothingPoolStatus(status bool) (api.CanSetSmoothingPoolRegistrationStatusResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-set-smoothing-pool-status %t", status))
	if err != nil {
		return api.CanSetSmoothingPoolRegistrationStatusResponse{}, fmt.Errorf("Could not get can-set-smoothing-pool-status: %w", err)
	}
	var response api.CanSetSmoothingPoolRegistrationStatusResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanSetSmoothingPoolRegistrationStatusResponse{}, fmt.Errorf("Could not decode can-set-smoothing-pool-status response: %w", err)
	}
	if response.Error != "" {
		return api.CanSetSmoothingPoolRegistrationStatusResponse{}, fmt.Errorf("Could not get can-set-smoothing-pool-status: %s", response.Error)
	}
	return response, nil
}

// Sets the node's Smoothing Pool opt-in status
func (c *Client) NodeSetSmoothingPoolStatus(status bool) (api.SetSmoothingPoolRegistrationStatusResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node set-smoothing-pool-status %t", status))
	if err != nil {
		return api.SetSmoothingPoolRegistrationStatusResponse{}, fmt.Errorf("Could not set smoothing pool status: %w", err)
	}
	var response api.SetSmoothingPoolRegistrationStatusResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SetSmoothingPoolRegistrationStatusResponse{}, fmt.Errorf("Could not decode set-smoothing-pool-status response: %w", err)
	}
	if response.Error != "" {
		return api.SetSmoothingPoolRegistrationStatusResponse{}, fmt.Errorf("Could not set smoothing pool status: %s", response.Error)
	}
	return response, nil
}

func (c *Client) ResolveEnsName(name string) (api.ResolveEnsNameResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node resolve-ens-name %s", name))
	if err != nil {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("Could not resolve ENS name: %w", err)
	}
	var response api.ResolveEnsNameResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("Could not decode resolve-ens-name: %w", err)
	}
	if response.Error != "" {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("Could not resolve ENS name: %s", response.Error)
	}
	return response, nil
}
func (c *Client) ReverseResolveEnsName(name string) (api.ResolveEnsNameResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node reverse-resolve-ens-name %s", name))
	if err != nil {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("Could not reverse resolve ENS name: %w", err)
	}
	var response api.ResolveEnsNameResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("Could not decode reverse-resolve-ens-name: %w", err)
	}
	if response.Error != "" {
		return api.ResolveEnsNameResponse{}, fmt.Errorf("Could not reverse resolve ENS name: %s", response.Error)
	}
	return response, nil
}

// Use the node private key to sign an arbitrary message
func (c *Client) SignMessage(message string) (api.NodeSignResponse, error) {
	responseBytes, err := c.callAPI("node sign-message", message)
	if err != nil {
		return api.NodeSignResponse{}, fmt.Errorf("Could not sign message: %w", err)
	}

	var response api.NodeSignResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSignResponse{}, fmt.Errorf("Could not decode node sign response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSignResponse{}, fmt.Errorf("Could not sign message: %s", response.Error)
	}
	return response, nil
}
