package codec

import (
	"github.com/shapeshift/cosmos-sdk/codec"
	cryptocodec "github.com/shapeshift/cosmos-sdk/crypto/codec"
	sdk "github.com/shapeshift/cosmos-sdk/types"
)

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	cryptocodec.RegisterCrypto(Amino)
	codec.RegisterEvidences(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
