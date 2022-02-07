package source

import (
	charitytypes "github.com/encichain/enci/x/charity/types"
)

type Source interface {
	Params(height int64) (charitytypes.Params, error)
}
