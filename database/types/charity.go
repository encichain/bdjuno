package types

// CharityParamsRow represents a single row inside the charity_params table
type CharityParamsRow struct {
	OneRowID bool   `db:"one_row_id"`
	Params   string `db:"params"`
	Height   int64  `db:"height"`
}

// NewMintParamsRow builds a new MintParamsRow instance
func NewCharityParamsRow(
	params string, height int64,
) CharityParamsRow {
	return CharityParamsRow{
		OneRowID: true,
		Params:   params,
		Height:   height,
	}
}

// Equal reports whether a and b represent the same table rows.
func (a CharityParamsRow) Equal(b CharityParamsRow) bool {
	return a.Params == b.Params &&
		a.Height == b.Height
}
