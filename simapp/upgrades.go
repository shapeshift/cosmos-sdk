package simapp

import (
	"github.com/shapeshift/cosmos-sdk/baseapp"
	storetypes "github.com/shapeshift/cosmos-sdk/store/types"
	sdk "github.com/shapeshift/cosmos-sdk/types"
	"github.com/shapeshift/cosmos-sdk/types/module"
	authtypes "github.com/shapeshift/cosmos-sdk/x/auth/types"
	banktypes "github.com/shapeshift/cosmos-sdk/x/bank/types"
	consensustypes "github.com/shapeshift/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/shapeshift/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/shapeshift/cosmos-sdk/x/distribution/types"
	govtypes "github.com/shapeshift/cosmos-sdk/x/gov/types"
	govv1 "github.com/shapeshift/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/shapeshift/cosmos-sdk/x/mint/types"
	paramstypes "github.com/shapeshift/cosmos-sdk/x/params/types"
	slashingtypes "github.com/shapeshift/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/shapeshift/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/shapeshift/cosmos-sdk/x/upgrade/types"
)

// UpgradeName defines the on-chain upgrade name for the sample SimApp upgrade
// from v046 to v047.
//
// NOTE: This upgrade defines a reference implementation of what an upgrade
// could look like when an application is migrating from Cosmos SDK version
// v0.46.x to v0.47.x.
const UpgradeName = "v046-to-v047"

func (app SimApp) RegisterUpgradeHandlers() {
	// Set param key table for params module migration
	for _, subspace := range app.ParamsKeeper.GetSubspaces() {
		subspace := subspace

		var keyTable paramstypes.KeyTable
		switch subspace.Name() {
		case authtypes.ModuleName:
			keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
		case banktypes.ModuleName:
			keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
		case stakingtypes.ModuleName:
			keyTable = stakingtypes.ParamKeyTable() //nolint:staticcheck
		case minttypes.ModuleName:
			keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
		case distrtypes.ModuleName:
			keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
		case slashingtypes.ModuleName:
			keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
		case govtypes.ModuleName:
			keyTable = govv1.ParamKeyTable() //nolint:staticcheck
		case crisistypes.ModuleName:
			keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck
		}

		if !subspace.HasKeyTable() {
			subspace.WithKeyTable(keyTable)
		}
	}

	baseAppLegacySS := app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())

	app.UpgradeKeeper.SetUpgradeHandler(
		UpgradeName,
		func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			// Migrate Tendermint consensus parameters from x/params module to a dedicated x/consensus module.
			baseapp.MigrateParams(ctx, baseAppLegacySS, &app.ConsensusParamsKeeper)

			// Note: this migration is optional,
			// You can include x/gov proposal migration documented in [UPGRADING.md](https://github.com/shapeshift/cosmos-sdk/blob/main/UPGRADING.md)

			return app.ModuleManager.RunMigrations(ctx, app.Configurator(), fromVM)
		},
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				consensustypes.ModuleName,
				crisistypes.ModuleName,
			},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}
