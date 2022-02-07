package database

import (
	"encoding/json"
	"fmt"

	"github.com/forbole/bdjuno/v2/types"
)

func (db *Db) SaveCharityParams(params *types.CharityParams) error {
	bz, err := json.Marshal(params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling charity params: %s", err)
	}
	stmt := `
	INSERT INTO charity_params (params, height)
	VALUES ($1, $2)
	ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
        height = excluded.height
	WHERE charity_params.height <= excluded.height`

	_, err = db.Sql.Exec(stmt, string(bz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing charity params into DB: %s", err)
	}

	return nil
}
