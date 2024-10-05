package lahmacun

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"bilge/x/lahmacun/keeper"
	"bilge/x/lahmacun/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	// Initialize Lahmacun list
	for _, lahmacun := range genState.LahmacunList {
		if err := k.StoreLahmacun(ctx, lahmacun); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Export Lahmacun list
	req := &types.QueryLahmacunRequest{}
	lahmacunList, err := k.Lahmacun(ctx, req)
	if err != nil {
		panic(err)
	}
	genesis.LahmacunList = []*types.Lahmacun{lahmacunList.Lahmacun}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
