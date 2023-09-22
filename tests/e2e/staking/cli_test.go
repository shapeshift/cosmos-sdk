//go:build e2e
// +build e2e

package testutil

import (
	"testing"

	"github.com/shapeshift/cosmos-sdk/simapp"
	"github.com/shapeshift/cosmos-sdk/testutil/network"
	"github.com/stretchr/testify/suite"
)

func TestE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig(simapp.NewTestNetworkFixture)
	cfg.NumValidators = 2
	suite.Run(t, NewE2ETestSuite(cfg))
}
