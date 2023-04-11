package keeper

import (
	"context"
	"fmt"

	"github.com/Peixer/superchain/x/superchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePythonCode(goCtx context.Context, msg *types.MsgCreatePythonCode) (*types.MsgCreatePythonCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pythonCode = types.PythonCode{
		Creator: msg.Creator,
		URI:     msg.URI,
	}

	id := k.AppendPythonCode(
		ctx,
		pythonCode,
	)

	return &types.MsgCreatePythonCodeResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePythonCode(goCtx context.Context, msg *types.MsgUpdatePythonCode) (*types.MsgUpdatePythonCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pythonCode = types.PythonCode{
		Creator: msg.Creator,
		Id:      msg.Id,
		URI:     msg.URI,
	}

	// Checks that the element exists
	val, found := k.GetPythonCode(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPythonCode(ctx, pythonCode)

	return &types.MsgUpdatePythonCodeResponse{}, nil
}

func (k msgServer) DeletePythonCode(goCtx context.Context, msg *types.MsgDeletePythonCode) (*types.MsgDeletePythonCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetPythonCode(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePythonCode(ctx, msg.Id)

	return &types.MsgDeletePythonCodeResponse{}, nil
}
