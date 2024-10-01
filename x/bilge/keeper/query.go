package keeper

import (
	"bilge/x/bilge/types"
)

var _ types.QueryServer = Keeper{}
