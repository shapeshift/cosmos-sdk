package testutil

import (
	_ "github.com/shapeshift/cosmos-sdk/x/auth"
	_ "github.com/shapeshift/cosmos-sdk/x/auth/tx/config"
	_ "github.com/shapeshift/cosmos-sdk/x/bank"
	_ "github.com/shapeshift/cosmos-sdk/x/consensus"
	_ "github.com/shapeshift/cosmos-sdk/x/distribution"
	_ "github.com/shapeshift/cosmos-sdk/x/genutil"
	_ "github.com/shapeshift/cosmos-sdk/x/mint"
	_ "github.com/shapeshift/cosmos-sdk/x/params"
	_ "github.com/shapeshift/cosmos-sdk/x/slashing"
	_ "github.com/shapeshift/cosmos-sdk/x/staking"

	"github.com/shapeshift/cosmos-sdk/core/appconfig"
	authtypes "github.com/shapeshift/cosmos-sdk/x/auth/types"
	banktypes "github.com/shapeshift/cosmos-sdk/x/bank/types"
	consensustypes "github.com/shapeshift/cosmos-sdk/x/consensus/types"
	distrtypes "github.com/shapeshift/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/shapeshift/cosmos-sdk/x/genutil/types"
	minttypes "github.com/shapeshift/cosmos-sdk/x/mint/types"
	paramstypes "github.com/shapeshift/cosmos-sdk/x/params/types"
	slashingtypes "github.com/shapeshift/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/shapeshift/cosmos-sdk/x/staking/types"

	runtimev1alpha1 "github.com/shapeshift/cosmos-sdk/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "github.com/shapeshift/cosmos-sdk/api/cosmos/app/v1alpha1"
	authmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/auth/module/v1"
	bankmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/bank/module/v1"
	consensusmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/consensus/module/v1"
	distrmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/distribution/module/v1"
	genutilmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/genutil/module/v1"
	mintmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/mint/module/v1"
	paramsmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/params/module/v1"
	slashingmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/slashing/module/v1"
	stakingmodulev1 "github.com/shapeshift/cosmos-sdk/api/cosmos/staking/module/v1"
	txconfigv1 "github.com/shapeshift/cosmos-sdk/api/cosmos/tx/config/v1"
)

var AppConfig = appconfig.Compose(&appv1alpha1.Config{
	Modules: []*appv1alpha1.ModuleConfig{
		{
			Name: "runtime",
			Config: appconfig.WrapAny(&runtimev1alpha1.Module{
				AppName: "StakingApp",
				BeginBlockers: []string{
					minttypes.ModuleName,
					distrtypes.ModuleName,
					stakingtypes.ModuleName,
					authtypes.ModuleName,
					banktypes.ModuleName,
					genutiltypes.ModuleName,
					slashingtypes.ModuleName,
					paramstypes.ModuleName,
					consensustypes.ModuleName,
				},
				EndBlockers: []string{
					stakingtypes.ModuleName,
					authtypes.ModuleName,
					banktypes.ModuleName,
					genutiltypes.ModuleName,
					distrtypes.ModuleName,
					minttypes.ModuleName,
					slashingtypes.ModuleName,
					paramstypes.ModuleName,
					consensustypes.ModuleName,
				},
				InitGenesis: []string{
					authtypes.ModuleName,
					banktypes.ModuleName,
					distrtypes.ModuleName,
					stakingtypes.ModuleName,
					minttypes.ModuleName,
					slashingtypes.ModuleName,
					genutiltypes.ModuleName,
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
					{Account: distrtypes.ModuleName},
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
			Name:   slashingtypes.ModuleName,
			Config: appconfig.WrapAny(&slashingmodulev1.Module{}),
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
			Name:   minttypes.ModuleName,
			Config: appconfig.WrapAny(&mintmodulev1.Module{}),
		},
		{
			Name:   distrtypes.ModuleName,
			Config: appconfig.WrapAny(&distrmodulev1.Module{}),
		},
	},
})
