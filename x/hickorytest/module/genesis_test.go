package hickorytest_test

import (
	"testing"

	keepertest "HICKORY-TEST/testutil/keeper"
	"HICKORY-TEST/testutil/nullify"
	hickorytest "HICKORY-TEST/x/hickorytest/module"
	"HICKORY-TEST/x/hickorytest/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HickorytestKeeper(t)
	hickorytest.InitGenesis(ctx, k, genesisState)
	got := hickorytest.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
