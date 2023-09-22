package module

import (
	sdk "github.com/shapeshift/cosmos-sdk/types"
	"github.com/shapeshift/cosmos-sdk/x/group/keeper"
)

// EndBlocker called at every block, updates proposal's `FinalTallyResult` and
// prunes expired proposals.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	if err := k.TallyProposalsAtVPEnd(ctx); err != nil {
		panic(err)
	}

	if err := k.PruneProposals(ctx); err != nil {
		panic(err)
	}
}
