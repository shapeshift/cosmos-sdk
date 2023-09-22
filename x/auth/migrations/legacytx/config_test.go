package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/shapeshift/cosmos-sdk/codec"
	cryptoAmino "github.com/shapeshift/cosmos-sdk/crypto/codec"
	"github.com/shapeshift/cosmos-sdk/testutil/testdata"
	sdk "github.com/shapeshift/cosmos-sdk/types"
	"github.com/shapeshift/cosmos-sdk/x/auth/migrations/legacytx"
	txtestutil "github.com/shapeshift/cosmos-sdk/x/auth/tx/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, txtestutil.NewTxConfigTestSuite(txGen))
}
