package simulation

import (
	"math/rand"

	"github.com/Peixer/superchain/x/superchain/keeper"
	"github.com/Peixer/superchain/x/superchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRunCode(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRunCode{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RunCode simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RunCode simulation not implemented"), nil, nil
	}
}
