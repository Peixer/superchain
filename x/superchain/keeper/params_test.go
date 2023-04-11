package keeper_test

import (
	"testing"

	testkeeper "github.com/Peixer/superchain/testutil/keeper"
	"github.com/Peixer/superchain/x/superchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SuperchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
