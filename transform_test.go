package erttransform_test

import (
	"github.com/enaml-ops/enaml"
	pushappsmanager "github.com/enaml-ops/ert-transform/push-apps-manager"

	. "github.com/enaml-ops/ert-transform"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ERTFromOSS", func() {
	Describe("app-manager", func() {
		var (
			manifest *enaml.DeploymentManifest
		)
		BeforeEach(func() {
			manifest = enaml.NewDeploymentManifest(nil)
		})

		It("adds tags", func() {
			t := &ERTFromOSS{}
			Ω(t.Apply(manifest)).ShouldNot(HaveOccurred())
			ig := manifest.GetInstanceGroupByName(pushappsmanager.InstanceGroupName)
			ij := ig.GetJobByName(pushappsmanager.InstanceJobName)
			Ω(ig).ShouldNot(BeNil())
			Ω(ij).ShouldNot(BeNil())
		})
	})
})
