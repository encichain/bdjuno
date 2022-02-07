package types

import (
	charitytypes "github.com/encichain/enci/x/charity/types"
)

// CharityParams represents the x/charity parameters
type CharityParams struct {
	Params charitytypes.Params
	Height int64
}

// NewCharityParams() creates a new CharityParams object
func NewCharityParams(params charitytypes.Params, height int64) *CharityParams {
	return &CharityParams{
		Params: params,
		Height: height,
	}
}
