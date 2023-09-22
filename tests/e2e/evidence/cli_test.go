//go:build e2e
// +build e2e

package evidence

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/shapeshift/cosmos-sdk/simapp"
	"github.com/shapeshift/cosmos-sdk/testutil/network"
)

func TestE2ETestSuite(t *testing.T) {
	cfg := network.DefaultConfig(simapp.NewTestNetworkFixture)
	cfg.NumValidators = 1
	suite.Run(t, NewE2ETestSuite(cfg))
}
