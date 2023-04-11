package superchain

import (
	"math/rand"

	"github.com/Peixer/superchain/testutil/sample"
	superchainsimulation "github.com/Peixer/superchain/x/superchain/simulation"
	"github.com/Peixer/superchain/x/superchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = superchainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreatePythonCode = "op_weight_msg_python_code"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePythonCode int = 100

	opWeightMsgUpdatePythonCode = "op_weight_msg_python_code"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePythonCode int = 100

	opWeightMsgDeletePythonCode = "op_weight_msg_python_code"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePythonCode int = 100

	opWeightMsgRunCode = "op_weight_msg_run_code"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRunCode int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	superchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PythonCodeList: []types.PythonCode{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PythonCodeCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&superchainGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePythonCode int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePythonCode, &weightMsgCreatePythonCode, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePythonCode = defaultWeightMsgCreatePythonCode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePythonCode,
		superchainsimulation.SimulateMsgCreatePythonCode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePythonCode int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePythonCode, &weightMsgUpdatePythonCode, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePythonCode = defaultWeightMsgUpdatePythonCode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePythonCode,
		superchainsimulation.SimulateMsgUpdatePythonCode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePythonCode int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeletePythonCode, &weightMsgDeletePythonCode, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePythonCode = defaultWeightMsgDeletePythonCode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePythonCode,
		superchainsimulation.SimulateMsgDeletePythonCode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRunCode int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRunCode, &weightMsgRunCode, nil,
		func(_ *rand.Rand) {
			weightMsgRunCode = defaultWeightMsgRunCode
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRunCode,
		superchainsimulation.SimulateMsgRunCode(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
