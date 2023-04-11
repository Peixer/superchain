package keeper

import (
	"github.com/Peixer/superchain/x/superchain/types"
)

var _ types.QueryServer = Keeper{}
