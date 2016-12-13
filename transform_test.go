package erttransform_test

import (
	"github.com/enaml-ops/enaml"
	pushappsmanager "github.com/enaml-ops/ert-transform/push-apps-manager"

	. "github.com/enaml-ops/ert-transform"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ERTFromOSS", func() {

	var (
		manifest *enaml.DeploymentManifest
	)
	BeforeEach(func() {
		manifest = enaml.NewDeploymentManifest(nil)
		ig := manifest.GetInstanceGroupByName(pushappsmanager.InstanceGroupName)
		Ω(ig).Should(BeNil())
	})
	Describe("setup config", func() {
		Context("when there is a manifest object to grab config state from", func() {
			var controlAZs = []string{"z1", "z2"}
			var controlStemcell = "trusty"
			var controlNetworkName = "net1"
			var instanceGroup *enaml.InstanceGroup
			BeforeEach(func() {
				manifest.Stemcells = []enaml.Stemcell{
					enaml.Stemcell{
						Name: controlStemcell,
					},
				}
				manifest.AddInstanceGroup(&enaml.InstanceGroup{
					AZs: controlAZs,
					Networks: []enaml.Network{
						enaml.Network{Name: controlNetworkName},
					},
				})
				t := &ERTFromOSS{}
				Ω(t.Apply(manifest)).ShouldNot(HaveOccurred())
				instanceGroup = manifest.GetInstanceGroupByName(pushappsmanager.InstanceGroupName)
			})

			It("should set azs on app-manager instance groups", func() {
				Ω(instanceGroup.AZs).Should(Equal(controlAZs))
			})

			It("should set stemcell on app-manager instance groups", func() {
				Ω(instanceGroup.Stemcell).Should(Equal(controlStemcell))
			})

			It("should set network name on app-manager instance groups", func() {
				Ω(instanceGroup.Networks).Should(HaveLen(1))
				Ω(instanceGroup.Networks[0].Name).Should(Equal(controlNetworkName))
			})
		})
	})
	Describe("app-manager", func() {

		It("adds and instance group with a job for app-manager", func() {
			t := &ERTFromOSS{}
			Ω(t.Apply(manifest)).ShouldNot(HaveOccurred())
			ig := manifest.GetInstanceGroupByName(pushappsmanager.InstanceGroupName)
			ij := ig.GetJobByName(pushappsmanager.InstanceJobName)
			Ω(ig).ShouldNot(BeNil())
			Ω(ij).ShouldNot(BeNil())
		})
	})
})
