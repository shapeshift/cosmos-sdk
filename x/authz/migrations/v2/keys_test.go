package v2

import (
	"testing"

	"github.com/shapeshift/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/shapeshift/cosmos-sdk/types"
	"github.com/shapeshift/cosmos-sdk/types/address"
	bank "github.com/shapeshift/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
)

var (
	granter = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	grantee = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	msgType = bank.SendAuthorization{}.MsgTypeURL()
)

func TestGrantkey(t *testing.T) {
	require := require.New(t)
	key := GrantStoreKey(grantee, granter, msgType)
	require.Len(key, len(GrantPrefix)+len(address.MustLengthPrefix(grantee))+len(address.MustLengthPrefix(granter))+len([]byte(msgType)))

	granter1, grantee1, msgType1 := ParseGrantKey(key[1:])
	require.Equal(granter, granter1)
	require.Equal(grantee, grantee1)
	require.Equal(msgType1, msgType)
}
