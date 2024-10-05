package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "bilge/testutil/keeper"
	"bilge/x/lahmacun/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LahmacunKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
