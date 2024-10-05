package keeper

import (
	"context"

	"bilge/x/lahmacun/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// CreateLahmacun handles the creation of a Lahmacun.
func (k msgServer) CreateLahmacun(goCtx context.Context, msg *types.MsgCreateLahmacun) (*types.MsgCreateLahmacunResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	lahmacun := types.Lahmacun{
		Id:     msg.Id,
		Heat:   msg.Heat,
		Dough:  msg.Dough,
		Meat:   msg.Meat,
		Tomato: msg.Tomato,
		Onion:  msg.Onion,
	}

	if err := k.Keeper.StoreLahmacun(ctx, &lahmacun); err != nil {
		return nil, err
	}

	return &types.MsgCreateLahmacunResponse{}, nil
}

// UpdateLahmacun handles the update of a Lahmacun.
func (k msgServer) UpdateLahmacun(goCtx context.Context, msg *types.MsgUpdateLahmacun) (*types.MsgUpdateLahmacunResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	lahmacun := types.Lahmacun{
		Id:     msg.Id,
		Heat:   msg.Heat,
		Dough:  msg.Dough,
		Meat:   msg.Meat,
		Tomato: msg.Tomato,
		Onion:  msg.Onion,
	}

	if err := k.Keeper.StoreLahmacun(ctx, &lahmacun); err != nil {
		return nil, err
	}

	return &types.MsgUpdateLahmacunResponse{}, nil
}

var _ types.MsgServer = msgServer{}
