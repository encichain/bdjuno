package remote

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/forbole/juno/v2/node/remote"

	distrsource "github.com/forbole/bdjuno/v2/modules/distribution/source"
)

var (
	_ distrsource.Source = &Source{}
)

// Source implements distrsource.Source querying the data from a remote node
type Source struct {
	*remote.Source
	distrClient distrtypes.QueryClient
}

// NewSource returns a new Source instace
func NewSource(source *remote.Source, distrClient distrtypes.QueryClient) *Source {
	return &Source{
		Source:      source,
		distrClient: distrClient,
	}
}

// ValidatorCommission implements distrsource.Source
func (s Source) ValidatorCommission(valOperAddr string, height int64) (sdk.DecCoins, error) {
	res, err := s.distrClient.ValidatorCommission(
		remote.GetHeightRequestContext(s.Ctx, height),
		&distrtypes.QueryValidatorCommissionRequest{ValidatorAddress: valOperAddr},
	)
	if err != nil {
		return nil, err
	}

	return res.Commission.Commission, nil
}

// DelegatorTotalRewards implements distrsource.Source
func (s Source) DelegatorTotalRewards(delegator string, height int64) ([]distrtypes.DelegationDelegatorReward, error) {
	res, err := s.distrClient.DelegationTotalRewards(
		remote.GetHeightRequestContext(s.Ctx, height),
		&distrtypes.QueryDelegationTotalRewardsRequest{DelegatorAddress: delegator},
	)
	if err != nil {
		return nil, fmt.Errorf("error while getting delegation total rewards for for delegator %s at height %v: %s", delegator, height, err)
	}

	return res.Rewards, nil
}

// DelegatorWithdrawAddress implements distrsource.Source
func (s Source) DelegatorWithdrawAddress(delegator string, height int64) (string, error) {
	res, err := s.distrClient.DelegatorWithdrawAddress(
		remote.GetHeightRequestContext(s.Ctx, height),
		&distrtypes.QueryDelegatorWithdrawAddressRequest{DelegatorAddress: delegator},
	)
	if err != nil {
		return "", err
	}

	return res.WithdrawAddress, nil
}

// CommunityPool implements distrsource.Source
func (s Source) CommunityPool(height int64) (sdk.DecCoins, error) {
	res, err := s.distrClient.CommunityPool(
		remote.GetHeightRequestContext(s.Ctx, height),
		&distrtypes.QueryCommunityPoolRequest{},
	)
	if err != nil {
		return nil, err
	}

	return res.Pool, nil
}

// Params implements distrsource.Source
func (s Source) Params(height int64) (distrtypes.Params, error) {
	res, err := s.distrClient.Params(
		remote.GetHeightRequestContext(s.Ctx, height),
		&distrtypes.QueryParamsRequest{},
	)
	if err != nil {
		return distrtypes.Params{}, err
	}

	return res.Params, nil
}
