package nft

import (
	types "github.com/shapeshift/cosmos-sdk/codec/types"
	sdk "github.com/shapeshift/cosmos-sdk/types"
	"github.com/shapeshift/cosmos-sdk/types/msgservice"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
