package keeper

import (
	"bilge/x/lahmacun/types"
)

var _ types.QueryServer = Keeper{}
