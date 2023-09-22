package authz

import (
	"github.com/shapeshift/cosmos-sdk/client"
	"github.com/shapeshift/cosmos-sdk/testutil"
	clitestutil "github.com/shapeshift/cosmos-sdk/testutil/cli"
	"github.com/shapeshift/cosmos-sdk/x/authz/client/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
