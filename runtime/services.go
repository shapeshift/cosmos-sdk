package runtime

import (
	appv1alpha1 "github.com/shapeshift/cosmos-sdk/api/cosmos/app/v1alpha1"
	autocliv1 "github.com/shapeshift/cosmos-sdk/api/cosmos/autocli/v1"
	reflectionv1 "github.com/shapeshift/cosmos-sdk/api/cosmos/reflection/v1"

	"github.com/shapeshift/cosmos-sdk/runtime/services"
)

func (a *App) registerRuntimeServices() error {
	appv1alpha1.RegisterQueryServer(a.GRPCQueryRouter(), services.NewAppQueryService(a.appConfig))
	autocliv1.RegisterQueryServer(a.GRPCQueryRouter(), services.NewAutoCLIQueryService(a.ModuleManager.Modules))

	reflectionSvc, err := services.NewReflectionService()
	if err != nil {
		return err
	}
	reflectionv1.RegisterReflectionServiceServer(a.GRPCQueryRouter(), reflectionSvc)

	return nil
}
