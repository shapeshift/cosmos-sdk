package testutil

import (
	"context"
	"fmt"

	tmcfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/cli"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/spf13/viper"

	"github.com/shapeshift/cosmos-sdk/client"
	"github.com/shapeshift/cosmos-sdk/codec"
	"github.com/shapeshift/cosmos-sdk/server"
	"github.com/shapeshift/cosmos-sdk/testutil"
	"github.com/shapeshift/cosmos-sdk/types/module"
	genutilcli "github.com/shapeshift/cosmos-sdk/x/genutil/client/cli"
)

func ExecInitCmd(testMbm module.BasicManager, home string, cdc codec.Codec) error {
	logger := log.NewNopLogger()
	cfg, err := CreateDefaultTendermintConfig(home)
	if err != nil {
		return err
	}

	cmd := genutilcli.InitCmd(testMbm, home)
	serverCtx := server.NewContext(viper.New(), cfg, logger)
	clientCtx := client.Context{}.WithCodec(cdc).WithHomeDir(home)

	_, out := testutil.ApplyMockIO(cmd)
	clientCtx = clientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)

	cmd.SetArgs([]string{"appnode-test", fmt.Sprintf("--%s=%s", cli.HomeFlag, home)})

	return cmd.ExecuteContext(ctx)
}

func CreateDefaultTendermintConfig(rootDir string) (*tmcfg.Config, error) {
	conf := tmcfg.DefaultConfig()
	conf.SetRoot(rootDir)
	tmcfg.EnsureRoot(rootDir)

	if err := conf.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("error in config file: %v", err)
	}

	return conf, nil
}
