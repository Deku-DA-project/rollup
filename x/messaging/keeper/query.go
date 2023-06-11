package keeper

import (
	"messaging/x/messaging/types"
)

var _ types.QueryServer = Keeper{}
