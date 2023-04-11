package keeper_test

import (
	"testing"

	keepertest "github.com/Peixer/superchain/testutil/keeper"
	"github.com/Peixer/superchain/testutil/nullify"
	"github.com/Peixer/superchain/x/superchain/keeper"
	"github.com/Peixer/superchain/x/superchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNPythonCode(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PythonCode {
	items := make([]types.PythonCode, n)
	for i := range items {
		items[i].Id = keeper.AppendPythonCode(ctx, items[i])
	}
	return items
}

func TestPythonCodeGet(t *testing.T) {
	keeper, ctx := keepertest.SuperchainKeeper(t)
	items := createNPythonCode(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPythonCode(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPythonCodeRemove(t *testing.T) {
	keeper, ctx := keepertest.SuperchainKeeper(t)
	items := createNPythonCode(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePythonCode(ctx, item.Id)
		_, found := keeper.GetPythonCode(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPythonCodeGetAll(t *testing.T) {
	keeper, ctx := keepertest.SuperchainKeeper(t)
	items := createNPythonCode(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPythonCode(ctx)),
	)
}

func TestPythonCodeCount(t *testing.T) {
	keeper, ctx := keepertest.SuperchainKeeper(t)
	items := createNPythonCode(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPythonCodeCount(ctx))
}
