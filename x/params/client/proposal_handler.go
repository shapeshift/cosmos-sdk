package client

import (
	govclient "github.com/shapeshift/cosmos-sdk/x/gov/client"
	"github.com/shapeshift/cosmos-sdk/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
