package database_test

import (
	"encoding/json"

	charitytypes "github.com/encichain/enci/x/charity/types"

	"github.com/forbole/bdjuno/v2/types"

	dbtypes "github.com/forbole/bdjuno/v2/database/types"
)

func (suite *DbTestSuite) TestSaveCharityParams() {
	require := suite.Require()
	charityParams := charitytypes.DefaultParams()

	//bdjunoCharityParams := types.NewCharityParams(charityParams, 10)
	//require.Equal(1, bdjunoCharityParams)
	err := suite.database.SaveCharityParams(types.NewCharityParams(charityParams, 10))
	require.NoError(err)

	var rows []dbtypes.CharityParamsRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM charity_params`)
	require.NoError(err)
	require.Len(rows, 1)

	var storedParams charitytypes.Params
	err = json.Unmarshal([]byte(rows[0].Params), &storedParams)
	require.NoError(err)
	require.Equal(charityParams, storedParams)
	require.Equal(int64(10), rows[0].Height)
}
