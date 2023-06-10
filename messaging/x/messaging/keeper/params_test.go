package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "messaging/testutil/keeper"
	"messaging/x/messaging/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MessagingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
