package keeper

import (
	"bilge/x/lahmacun/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StoreLahmacun stores a Lahmacun object based on the LahmacunRequest.
func (k Keeper) StoreLahmacun(ctx sdk.Context, lahmacun *types.Lahmacun) error {
	// Validate the Lahmacun object
	if err := lahmacun.Validate(); err != nil {
		return err
	}

	store := k.storeService.OpenKVStore(ctx)
	bz, err := k.cdc.Marshal(lahmacun)
	if err != nil {
		return err
	}

	if err := store.Set([]byte(lahmacun.Id), bz); err != nil {
		return err
	}

	return nil
}
