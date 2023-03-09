package keeper

import (
	"github.com/tomahawk-network/tomahawk/v6/x/govshuttle/types"
)

var _ types.QueryServer = Keeper{}
