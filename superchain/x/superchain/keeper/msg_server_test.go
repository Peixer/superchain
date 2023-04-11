package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Peixer/superchain/testutil/keeper"
	"github.com/Peixer/superchain/x/superchain/keeper"
	"github.com/Peixer/superchain/x/superchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SuperchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
