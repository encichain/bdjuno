package remote

import (
	"github.com/forbole/juno/v2/node/remote"

	charitytypes "github.com/encichain/enci/x/charity/types"
	charitysource "github.com/forbole/bdjuno/v2/modules/charity/source"
)

var (
	_ charitysource.Source = &Source{}
)

// Source implements charitysource.Source using a remote node
type Source struct {
	*remote.Source
	querier charitytypes.QueryClient
}

// NewSource returns a new Source object
func NewSourceRemote(source *remote.Source, querier charitytypes.QueryClient) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// Params() implements charitysource.Source
func (s Source) Params(height int64) (charitytypes.Params, error) {
	res, err := s.querier.Params(remote.GetHeightRequestContext(s.Ctx, height), &charitytypes.QueryParamsRequest{})
	if err != nil {
		return charitytypes.Params{}, err
	}

	return res.Params, nil
}
