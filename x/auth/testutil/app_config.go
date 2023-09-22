package testutil

import (
	_ "github.com/shapeshift/cosmos-sdk/x/auth"
	_ "github.com/shapeshift/cosmos-sdk/x/auth/tx/config"
	_ "github.com/shapeshift/cosmos-sdk/x/auth/vesting"
	_ "github.com/shapeshift/cosmos-sdk/x/bank"
	_ "github.com/shapeshift/cosmos-sdk/x/consensus"
	_ "github.com/shapeshift/cosmos-sdk/x/feegrant/module"
	_ "github.com/shapeshift/cosmos-sdk/x/genutil"
	_ "github.com/shapeshift/cosmos-sdk/x/params"
	_ "github.com/shapeshift/cosmos-sdk/x/staking"

	"github.com/shapeshift/cosmos-sdk/core/appconfig"
	authtypes "github.com/shapeshift/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/shapeshift/cosmos-sdk/x/auth/vesting/types"
	banktypes "github.com/shapeshift/cosmos-sdk/x/bank/types"
	consensustypes "github.com/shapeshift/cosmos-sdk/x/consensus/types"
	"github.com/shapeshift/cosmos-sdk/x/feegrant"
	genutiltypes "github.com/shapeshift/cosmos-sdk/x/genutil/types"
	minttypes "github.com/shapeshift/cosmos-sdk/x/mint/types"
	paramstypes "github.com/shapeshift/cosmos-sdk/x/params/types"
	stakingtypes "github.com/shapeshift/cosmos-sdk/x/staking/types"

	runtimev1alpha1 "github.com/shapeshift/cosmos-sdk/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "github.com/shapeshift/cosmos-sdk/api/cosmos/app/v1alpha1"
	authmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/auth/module/v1"
	bankmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/bank/module/v1"
	consensusmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/consensus/module/v1"
	feegrantmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/feegrant/module/v1"
	genutilmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/genutil/module/v1"
	paramsmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/params/module/v1"
	stakingmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/staking/module/v1"
	txconfigv1 "github.com/shapeshift/cosmos-sdk/api/cosmos/tx/config/v1"
	vestingmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/vesting/module/v1"
)

var AppConfig = appconfig.Compose(&appv1alpha1.Config{
	Modules: []*appv1alpha1.ModuleConfig{
		{
			Name: "runtime",
			Config: appconfig.WrapAny(&runtimev1alpha1.Module{
				AppName: "AuthApp",
				BeginBlockers: []string{
					stakingtypes.ModuleName,
					authtypes.ModuleName,
					banktypes.ModuleName,
					genutiltypes.ModuleName,
					feegrant.ModuleName,
					paramstypes.ModuleName,
					consensustypes.ModuleName,
					vestingtypes.ModuleName,
				},
				EndBlockers: []string{
					stakingtypes.ModuleName,
					authtypes.ModuleName,
					banktypes.ModuleName,
					genutiltypes.ModuleName,
					feegrant.ModuleName,
					paramstypes.ModuleName,
					consensustypes.ModuleName,
					vestingtypes.ModuleName,
				},
				InitGenesis: []string{
					authtypes.ModuleName,
					banktypes.ModuleName,
					stakingtypes.ModuleName,
					genutiltypes.ModuleName,
					feegrant.ModuleName,
					paramstypes.ModuleName,
					consensustypes.ModuleName,
					vestingtypes.ModuleName,
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
					{Account: "multiple permissions account", Permissions: []string{authtypes.Minter, authtypes.Burner, stakingtypes.ModuleName}}, // dummy permissions
					{Account: "random permission", Permissions: []string{"random"}},
				},
			}),
		},
		{
			Name:   vestingtypes.ModuleName,
			Config: appconfig.WrapAny(&vestingmodulev1.Module{}),
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
			Name:   feegrant.ModuleName,
			Config: appconfig.WrapAny(&feegrantmodulev1.Module{}),
		},
	},
})
