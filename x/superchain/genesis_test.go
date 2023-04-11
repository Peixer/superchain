package superchain_test

import (
	"testing"

	keepertest "github.com/Peixer/superchain/testutil/keeper"
	"github.com/Peixer/superchain/testutil/nullify"
	"github.com/Peixer/superchain/x/superchain"
	"github.com/Peixer/superchain/x/superchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PythonCodeList: []types.PythonCode{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PythonCodeCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SuperchainKeeper(t)
	superchain.InitGenesis(ctx, *k, genesisState)
	got := superchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PythonCodeList, got.PythonCodeList)
	require.Equal(t, genesisState.PythonCodeCount, got.PythonCodeCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
