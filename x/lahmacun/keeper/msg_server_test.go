package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "bilge/testutil/keeper"
	"bilge/x/lahmacun/keeper"
	"bilge/x/lahmacun/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, sdk.Context) {
	k, ctx := keepertest.LahmacunKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
