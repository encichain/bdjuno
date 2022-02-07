package charity

import (
	"encoding/json"
	"fmt"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/forbole/bdjuno/v2/types"

	charitytypes "github.com/encichain/enci/x/charity/types"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements the module.GenesisModule interface
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "charity").Msg("parsing genesis")

	// Read the genesis state
	var genState charitytypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[charitytypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while unmarshaling charity genesis data: %s", err)
	}

	// Save the charity params to DB
	err = m.db.SaveCharityParams(types.NewCharityParams(genState.Params, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error storing genesis charity params to DB: %s", err)
	}
	return nil
}
