package auth_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	simtestutil "github.com/shapeshift/cosmos-sdk/testutil/sims"
	"github.com/shapeshift/cosmos-sdk/x/auth/keeper"
	"github.com/shapeshift/cosmos-sdk/x/auth/testutil"
	"github.com/shapeshift/cosmos-sdk/x/auth/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	var accountKeeper keeper.AccountKeeper
	app, err := simtestutil.SetupAtGenesis(testutil.AppConfig, &accountKeeper)
	require.NoError(t, err)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	acc := accountKeeper.GetAccount(ctx, types.NewModuleAddress(types.FeeCollectorName))
	require.NotNil(t, acc)
}
