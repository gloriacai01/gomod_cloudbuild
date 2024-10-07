package models

import (
	"errors"

	"go.viam.com/rdk/components/arm"
	"go.viam.com/utils/rpc"

	"context"

	pb "go.viam.com/api/component/arm/v1"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
)

var (
	Mymodel          = resource.NewModel("3531d2cd-ea2d-49e2-b526-975e21943ff6", "gomod_cloudbuild", "mymodel")
	errUnimplemented = errors.New("unimplemented")
)

func init() {
	resource.RegisterComponent(arm.API, Mymodel,
		resource.Registration[arm.Arm, *Config]{
			Constructor: newGomodCloudbuildMymodel,
		},
	)
}

type Config struct {
	// Put config attributes here

	/* if your model  does not need a config,
	   replace *Config in the init function with resource.NoNativeConfig */

	/* Uncomment this if your model does not need to be validated
	   and has no implicit dependecies. */
	// resource.TriviallyValidateConfig

}

func (cfg *Config) Validate(path string) ([]string, error) {
	// Add config validation code here
	return nil, nil
}

type gomodCloudbuildMymodel struct {
	name resource.Name

	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()

	/* Uncomment this if your model does not need to reconfigure. */
	// resource.TriviallyReconfigurable

	// Uncomment this if the model does not have any goroutines that
	// need to be shut down while closing.
	// resource.TriviallyCloseable

}

func newGomodCloudbuildMymodel(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (arm.Arm, error) {
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &gomodCloudbuildMymodel{
		name:       rawConf.ResourceName(),
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}
	return s, nil
}

func (s *gomodCloudbuildMymodel) Name() resource.Name {
	return s.name
}

func (s *gomodCloudbuildMymodel) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
	// Put reconfigure code here
	return errUnimplemented
}

func (s *gomodCloudbuildMymodel) NewClientFromConn(ctx context.Context, conn rpc.ClientConn, remoteName string, name resource.Name, logger logging.Logger) (arm.Arm, error) {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) EndPosition(ctx context.Context, extra map[string]interface{}) (spatialmath.Pose, error) {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) MoveToPosition(ctx context.Context, pose spatialmath.Pose, extra map[string]interface{}) error {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) MoveToJointPositions(ctx context.Context, positions *pb.JointPositions, extra map[string]interface{}) error {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) JointPositions(ctx context.Context, extra map[string]interface{}) (*pb.JointPositions, error) {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) Stop(ctx context.Context, extra map[string]interface{}) error {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) ModelFrame() referenceframe.Model {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) CurrentInputs(ctx context.Context) ([]referenceframe.Input, error) {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) GoToInputs(ctx context.Context, inputSteps ...[]referenceframe.Input) error {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) IsMoving(ctx context.Context) (bool, error) {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) Geometries(ctx context.Context, extra map[string]interface{}) ([]spatialmath.Geometry, error) {
	panic("not implemented")
}

func (s *gomodCloudbuildMymodel) Close(context.Context) error {
	s.cancelFunc()
	return nil
}
