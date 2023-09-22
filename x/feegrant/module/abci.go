package module

import (
	sdk "github.com/shapeshift/cosmos-sdk/types"
	"github.com/shapeshift/cosmos-sdk/x/feegrant/keeper"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.RemoveExpiredAllowances(ctx)
}
