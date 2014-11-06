package vm_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	bmconfig "github.com/cloudfoundry/bosh-micro-cli/config"
	bmdepl "github.com/cloudfoundry/bosh-micro-cli/deployment"

	fakesys "github.com/cloudfoundry/bosh-agent/system/fakes"
	fakebmcloud "github.com/cloudfoundry/bosh-micro-cli/cloud/fakes"

	. "github.com/cloudfoundry/bosh-micro-cli/deployer/vm"
)

var _ = Describe("Manager", func() {
	Describe("CreateVM", func() {
		var (
			fakeCloud               *fakebmcloud.FakeCloud
			manager                 Manager
			expectedNetworksSpec    map[string]interface{}
			expectedCloudProperties map[string]interface{}
			expectedEnv             map[string]interface{}
			deployment              bmdepl.Deployment
			configService           bmconfig.DeploymentConfigService
			fs                      *fakesys.FakeFileSystem
		)

		BeforeEach(func() {
			logger := boshlog.NewLogger(boshlog.LevelNone)
			fs = fakesys.NewFakeFileSystem()
			configService = bmconfig.NewFileSystemDeploymentConfigService("/fake/path", fs, logger)
			fakeCloud = fakebmcloud.NewFakeCloud()
			manager = NewManagerFactory(configService, logger).NewManager(fakeCloud)
			fakeCloud.CreateVMCID = "fake-vm-cid"
			expectedNetworksSpec = map[string]interface{}{
				"fake-network-name": map[string]interface{}{
					"type":             "dynamic",
					"ip":               "fake-micro-ip",
					"cloud_properties": map[string]interface{}{},
				},
			}
			expectedCloudProperties = map[string]interface{}{
				"fake-cloud-property-key": "fake-cloud-property-value",
			}
			expectedEnv = map[string]interface{}{
				"fake-env-key": "fake-env-value",
			}
			deployment = bmdepl.Deployment{
				Name: "fake-deployment",
				Networks: []bmdepl.Network{
					{
						Name: "fake-network-name",
						Type: "dynamic",
					},
				},
				ResourcePools: []bmdepl.ResourcePool{
					{
						Name: "fake-resource-pool-name",
						RawCloudProperties: map[interface{}]interface{}{
							"fake-cloud-property-key": "fake-cloud-property-value",
						},
						RawEnv: map[interface{}]interface{}{
							"fake-env-key": "fake-env-value",
						},
					},
				},
				Jobs: []bmdepl.Job{
					{
						Name: "fake-job",
						Networks: []bmdepl.JobNetwork{
							{
								Name:      "fake-network-name",
								StaticIPs: []string{"fake-micro-ip"},
							},
						},
					},
				},
			}
		})

		It("creates a VM", func() {
			vm, err := manager.Create("fake-stemcell-cid", deployment)
			Expect(err).ToNot(HaveOccurred())
			Expect(vm).To(Equal(VM{
				CID: "fake-vm-cid",
			}))
			Expect(fakeCloud.CreateVMInput).To(Equal(
				fakebmcloud.CreateVMInput{
					StemcellCID:     "fake-stemcell-cid",
					CloudProperties: expectedCloudProperties,
					NetworksSpec:    expectedNetworksSpec,
					Env:             expectedEnv,
				},
			))
		})

		It("saves the vm record using the config service", func() {
			_, err := manager.Create("fake-stemcell-cid", deployment)
			Expect(err).ToNot(HaveOccurred())

			deploymentConfig, err := configService.Load()
			Expect(err).ToNot(HaveOccurred())

			expectedConfig := bmconfig.DeploymentConfig{
				VMCID: "fake-vm-cid",
			}
			Expect(deploymentConfig).To(Equal(expectedConfig))
		})

		Context("when creating the vm fails", func() {
			BeforeEach(func() {
				fakeCloud.CreateVMErr = errors.New("fake-create-error")
			})

			It("returns an error", func() {
				_, err := manager.Create("fake-stemcell-cid", deployment)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("fake-create-error"))
			})
		})
	})
})