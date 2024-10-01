package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "bilge/testutil/keeper"
	"bilge/x/bilge/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BilgeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
