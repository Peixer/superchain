package keeper

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/Peixer/superchain/x/superchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RunCode(goCtx context.Context, msg *types.MsgRunCode) (*types.MsgRunCodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// run python code
	// Get the ipfs
	val, found := k.GetPythonCode(ctx, uint64(msg.Id))
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	filepath := val.URI + ".py"

	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		downloadFile(filepath, "https://gateway.pinata.cloud/ipfs/"+val.URI)
	}

	cmd := exec.Command("python3", filepath)
	stdout, err := cmd.Output()

	messageReturn := ""
	if err != nil {
		messageReturn = err.Error()
	} else {
		messageReturn = string(stdout)
	}

	newIndex := strconv.FormatUint(uint64(msg.Id), 10)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.GameCreatedEventType,
			sdk.NewAttribute(types.GameCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.GameCreatedEventGameIndex, newIndex),
			sdk.NewAttribute(types.GameCreatedEventRed, string(messageReturn)),
		),
	)

	return &types.MsgRunCodeResponse{}, nil
}

func downloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
