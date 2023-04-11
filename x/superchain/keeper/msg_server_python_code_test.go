package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/Peixer/superchain/x/superchain/types"
)

func TestPythonCodeMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreatePythonCode(ctx, &types.MsgCreatePythonCode{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestPythonCodeMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdatePythonCode
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdatePythonCode{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdatePythonCode{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdatePythonCode{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreatePythonCode(ctx, &types.MsgCreatePythonCode{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdatePythonCode(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPythonCodeMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeletePythonCode
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeletePythonCode{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeletePythonCode{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeletePythonCode{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreatePythonCode(ctx, &types.MsgCreatePythonCode{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeletePythonCode(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
