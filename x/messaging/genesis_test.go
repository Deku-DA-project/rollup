package messaging_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "messaging/testutil/keeper"
	"messaging/testutil/nullify"
	"messaging/x/messaging"
	"messaging/x/messaging/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MessagingKeeper(t)
	messaging.InitGenesis(ctx, *k, genesisState)
	got := messaging.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
