package charity

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/bdjuno/v2/types"
)

// UpdateParams gets the updated params and stores them inside the database
func (m *Module) UpdateParams(height int64) error {
	log.Debug().Str("module", "charity").Int64("height", height).
		Msg("updating params")

	params, err := m.source.Params(height)
	if err != nil {
		return fmt.Errorf("error while getting params: %s", err)
	}

	return m.db.SaveCharityParams(types.NewCharityParams(params, height))

}

func (m *Module) PeriodicUpdateParams() error {
	height, err := m.db.LastBlockHeight()
	if err != nil {
		return fmt.Errorf("error while getting last block height: %s", err)
	}
	return m.UpdateParams(height)
}
