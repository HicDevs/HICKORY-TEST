package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "HICKORY-TEST/testutil/keeper"
	"HICKORY-TEST/x/hickorytest/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.HickorytestKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
