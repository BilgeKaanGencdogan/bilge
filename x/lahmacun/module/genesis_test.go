package lahmacun_test

import (
	"testing"

	keepertest "bilge/testutil/keeper"
	"bilge/testutil/nullify"
	lahmacun "bilge/x/lahmacun/module"
	"bilge/x/lahmacun/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LahmacunKeeper(t)
	lahmacun.InitGenesis(ctx, k, genesisState)
	got := lahmacun.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
