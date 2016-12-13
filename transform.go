package erttransform

import (
	"flag"

	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/ert-transform/config"
	pushappsmanager "github.com/enaml-ops/ert-transform/push-apps-manager"
	"github.com/enaml-ops/omg-transform/manifest"
)

// ERTFromOSS is a transformation that changes which network
// an instance group is placed in.
type ERTFromOSS struct {
}

//Apply -- implements transformation interface: will get called to create the transformmed manifest
func (s *ERTFromOSS) Apply(dm *enaml.DeploymentManifest) error {
	var err error
	c := &config.Config{}
	igc := pushappsmanager.NewPushAppsManager(c)
	ig := igc.ToInstanceGroup()
	dm.AddInstanceGroup(ig)
	return err
}

func (s *ERTFromOSS) flagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("ert-from-oss", flag.ContinueOnError)
	//fs.StringVar(&s.InstanceGroup, "instance-group", "", "name of the instance group")
	return fs
}

// ERTFromOSSTransformation is a TransformationBuilder that builds the
// 'ert-from-oss' transformation.
func ERTFromOSSTransformation(args []string) (manifest.Transformation, error) {
	s := &ERTFromOSS{}
	fs := s.flagSet()
	err := fs.Parse(args)
	return s, err
}
