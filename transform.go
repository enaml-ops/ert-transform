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
	c := new(config.Config)
	addAZs(dm, c)
	addNetwork(dm, c)
	addStemcell(dm, c)
	addAppsManager(dm, c)
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

func addStemcell(dm *enaml.DeploymentManifest, c *config.Config) {
	if len(dm.Stemcells) > 0 {
		c.StemcellName = dm.Stemcells[0].Name
	}
}

func addNetwork(dm *enaml.DeploymentManifest, c *config.Config) {
	if len(dm.InstanceGroups) > 0 && len(dm.InstanceGroups[0].Networks) > 0 {
		c.NetworkName = dm.InstanceGroups[0].Networks[0].Name
	}
}

func addAZs(dm *enaml.DeploymentManifest, c *config.Config) {
	if len(dm.InstanceGroups) > 0 {
		c.AZs = dm.InstanceGroups[0].AZs
	}
}

func addAppsManager(dm *enaml.DeploymentManifest, c *config.Config) {
	igc := pushappsmanager.NewPushAppsManager(c)
	ig := igc.ToInstanceGroup()
	dm.AddInstanceGroup(ig)
	dm.AddRelease(enaml.Release{
		Name:    pushappsmanager.PushAppsReleaseName,
		Version: pushappsmanager.PushAppsReleaseVersion,
	})
}
