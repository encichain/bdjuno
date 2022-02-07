package local

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v2/node/local"

	charitytypes "github.com/encichain/enci/x/charity/types"
	charitysource "github.com/forbole/bdjuno/v2/modules/charity/source"
)

var (
	_ charitysource.Source = &Source{}
)

// Source implements charitysource.Source using a remote node
type Source struct {
	*local.Source
	querier charitytypes.QueryClient
}

// NewSource returns a new Source object
func NewSourceLocal(source *local.Source, querier charitytypes.QueryClient) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// Params() implements charitysource.Source
func (s Source) Params(height int64) (charitytypes.Params, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return charitytypes.Params{}, fmt.Errorf("error loading height: %s", err)
	}

	res, err := s.querier.Params(sdk.WrapSDKContext(ctx), &charitytypes.QueryParamsRequest{})
	if err != nil {
		return charitytypes.Params{}, err
	}

	return res.Params, nil
}
