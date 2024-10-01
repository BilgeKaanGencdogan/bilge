package bilge_test

import (
	"testing"

	keepertest "bilge/testutil/keeper"
	"bilge/testutil/nullify"
	bilge "bilge/x/bilge/module"
	"bilge/x/bilge/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BilgeKeeper(t)
	bilge.InitGenesis(ctx, k, genesisState)
	got := bilge.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
