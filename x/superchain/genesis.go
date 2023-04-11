package superchain

import (
	"github.com/Peixer/superchain/x/superchain/keeper"
	"github.com/Peixer/superchain/x/superchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the pythonCode
	for _, elem := range genState.PythonCodeList {
		k.SetPythonCode(ctx, elem)
	}

	// Set pythonCode count
	k.SetPythonCodeCount(ctx, genState.PythonCodeCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PythonCodeList = k.GetAllPythonCode(ctx)
	genesis.PythonCodeCount = k.GetPythonCodeCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
