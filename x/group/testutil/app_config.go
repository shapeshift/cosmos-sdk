package testutil

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"

	_ "github.com/shapeshift/cosmos-sdk/x/auth"
	_ "github.com/shapeshift/cosmos-sdk/x/auth/tx/config"
	_ "github.com/shapeshift/cosmos-sdk/x/authz"
	_ "github.com/shapeshift/cosmos-sdk/x/bank"
	_ "github.com/shapeshift/cosmos-sdk/x/consensus"
	_ "github.com/shapeshift/cosmos-sdk/x/genutil"
	_ "github.com/shapeshift/cosmos-sdk/x/group/module"
	_ "github.com/shapeshift/cosmos-sdk/x/mint"
	_ "github.com/shapeshift/cosmos-sdk/x/params"
	_ "github.com/shapeshift/cosmos-sdk/x/staking"

	"cosmossdk.io/core/appconfig"
	authtypes "github.com/shapeshift/cosmos-sdk/x/auth/types"
	banktypes "github.com/shapeshift/cosmos-sdk/x/bank/types"
	consensustypes "github.com/shapeshift/cosmos-sdk/x/consensus/types"
	genutiltypes "github.com/shapeshift/cosmos-sdk/x/genutil/types"
	"github.com/shapeshift/cosmos-sdk/x/group"
	minttypes "github.com/shapeshift/cosmos-sdk/x/mint/types"
	paramstypes "github.com/shapeshift/cosmos-sdk/x/params/types"
	stakingtypes "github.com/shapeshift/cosmos-sdk/x/staking/types"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	groupmodulev1 "cosmossdk.io/api/cosmos/group/module/v1"
	paramsmodulev1 "cosmossdk.io/api/cosmos/params/module/v1"
	stakingmodulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
)

var AppConfig = appconfig.Compose(&appv1alpha1.Config{
	Modules: []*appv1alpha1.ModuleConfig{
		{
			Name: "runtime",
			Config: appconfig.WrapAny(&runtimev1alpha1.Module{
				AppName: "GroupApp",
				BeginBlockers: []string{
					minttypes.ModuleName,
					stakingtypes.ModuleName,
					authtypes.ModuleName,
					banktypes.ModuleName,
					genutiltypes.ModuleName,
					group.ModuleName,
					paramstypes.ModuleName,
					consensustypes.ModuleName,
				},
				EndBlockers: []string{
					stakingtypes.ModuleName,
					authtypes.ModuleName,
					banktypes.ModuleName,
					minttypes.ModuleName,
					genutiltypes.ModuleName,
					group.ModuleName,
					paramstypes.ModuleName,
					consensustypes.ModuleName,
				},
				InitGenesis: []string{
					authtypes.ModuleName,
					banktypes.ModuleName,
					stakingtypes.ModuleName,
					minttypes.ModuleName,
					genutiltypes.ModuleName,
					group.ModuleName,
					paramstypes.ModuleName,
					consensustypes.ModuleName,
				},
			}),
		},
		{
			Name: authtypes.ModuleName,
			Config: appconfig.WrapAny(&authmodulev1.Module{
				Bech32Prefix: "cosmos",
				ModuleAccountPermissions: []*authmodulev1.ModuleAccountPermission{
					{Account: authtypes.FeeCollectorName},
					{Account: minttypes.ModuleName, Permissions: []string{authtypes.Minter}},
					{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
					{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
				},
			}),
		},
		{
			Name:   banktypes.ModuleName,
			Config: appconfig.WrapAny(&bankmodulev1.Module{}),
		},
		{
			Name:   stakingtypes.ModuleName,
			Config: appconfig.WrapAny(&stakingmodulev1.Module{}),
		},
		{
			Name:   paramstypes.ModuleName,
			Config: appconfig.WrapAny(&paramsmodulev1.Module{}),
		},
		{
			Name:   consensustypes.ModuleName,
			Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
		},
		{
			Name:   "tx",
			Config: appconfig.WrapAny(&txconfigv1.Config{}),
		},
		{
			Name:   genutiltypes.ModuleName,
			Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
		},
		{
			Name: group.ModuleName,
			Config: appconfig.WrapAny(&groupmodulev1.Module{
				MaxExecutionPeriod: durationpb.New(time.Second * 1209600),
				MaxMetadataLen:     255,
			}),
		},
	},
})
