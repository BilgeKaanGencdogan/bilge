package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bilge/x/lahmacun/types"
)

var _ types.QueryServer = Keeper{}

// Lahmacun implements the QueryServer interface for querying Lahmacuns.
func (k Keeper) Lahmacun(goCtx context.Context, req *types.QueryLahmacunRequest) (*types.QueryLahmacunResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request: id is required")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get([]byte(req.Id))
	if err != nil {
		return nil, err
	}
	if bz == nil {
		return nil, fmt.Errorf("lahmacun not found")
	}

	var lahmacun types.Lahmacun
	if err := k.cdc.Unmarshal(bz, &lahmacun); err != nil {
		return nil, err
	}

	return &types.QueryLahmacunResponse{Lahmacun: &lahmacun}, nil
}
