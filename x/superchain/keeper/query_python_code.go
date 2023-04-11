package keeper

import (
	"context"

	"github.com/Peixer/superchain/x/superchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PythonCodeAll(goCtx context.Context, req *types.QueryAllPythonCodeRequest) (*types.QueryAllPythonCodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pythonCodes []types.PythonCode
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	pythonCodeStore := prefix.NewStore(store, types.KeyPrefix(types.PythonCodeKey))

	pageRes, err := query.Paginate(pythonCodeStore, req.Pagination, func(key []byte, value []byte) error {
		var pythonCode types.PythonCode
		if err := k.cdc.Unmarshal(value, &pythonCode); err != nil {
			return err
		}

		pythonCodes = append(pythonCodes, pythonCode)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPythonCodeResponse{PythonCode: pythonCodes, Pagination: pageRes}, nil
}

func (k Keeper) PythonCode(goCtx context.Context, req *types.QueryGetPythonCodeRequest) (*types.QueryGetPythonCodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	pythonCode, found := k.GetPythonCode(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPythonCodeResponse{PythonCode: pythonCode}, nil
}
