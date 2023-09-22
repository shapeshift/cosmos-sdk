package keeper

import (
	"context"

	tmtypes "github.com/cometbft/cometbft/types"

	sdk "github.com/shapeshift/cosmos-sdk/types"
	sdkerrors "github.com/shapeshift/cosmos-sdk/types/errors"
	"github.com/shapeshift/cosmos-sdk/x/consensus/types"
	govtypes "github.com/shapeshift/cosmos-sdk/x/gov/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.GetAuthority() != req.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	consensusParams := req.ToProtoConsensusParams()
	if err := tmtypes.ConsensusParamsFromProto(consensusParams).ValidateBasic(); err != nil {
		return nil, err
	}

	k.Set(ctx, &consensusParams)

	return &types.MsgUpdateParamsResponse{}, nil
}
