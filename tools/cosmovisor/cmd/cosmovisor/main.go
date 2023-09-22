package main

import (
	"context"
	"os"

	"github.com/shapeshift/cosmos-sdk/tools/cosmovisor"
	cverrors "github.com/shapeshift/cosmos-sdk/tools/cosmovisor/errors"
)

func main() {
	logger := cosmovisor.NewLogger()
	ctx := context.WithValue(context.Background(), cosmovisor.LoggerKey, logger)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		cverrors.LogErrors(logger, "", err)
		os.Exit(1)
	}
}
