package charity

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v2/modules"

	"github.com/forbole/bdjuno/v2/database"
	charitysource "github.com/forbole/bdjuno/v2/modules/charity/source"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the charity module/database
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source charitysource.Source
}

// NewModule returns a new module object
func NewModule(cdc codec.Codec, db *database.Db, source charitysource.Source) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		source: source,
	}
}

// Name() implements the module.Module interface
func (m Module) Name() string {
	return "charity"
}
