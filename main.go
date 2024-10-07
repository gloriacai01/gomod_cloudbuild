package main

import (
	"context"
	"gomod_cloudbuild/models"

	"go.viam.com/rdk/components/arm"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, module.NewLoggerFromArgs("gomod_cloudbuild"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	gomodCloudbuild, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	if err = gomodCloudbuild.AddModelFromRegistry(ctx, arm.API, models.Mymodel); err != nil {
		return err
	}

	err = gomodCloudbuild.Start(ctx)
	defer gomodCloudbuild.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
