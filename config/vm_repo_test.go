package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	fakesys "github.com/cloudfoundry/bosh-agent/system/fakes"

	. "github.com/cloudfoundry/bosh-micro-cli/config"
)

var _ = Describe("VMRepo", func() {
	var (
		repo          VMRepo
		configService DeploymentConfigService
		fs            *fakesys.FakeFileSystem
	)

	BeforeEach(func() {
		logger := boshlog.NewLogger(boshlog.LevelNone)
		fs = fakesys.NewFakeFileSystem()
		configService = NewFileSystemDeploymentConfigService("/fake/path", fs, logger)
		repo = NewVMRepo(configService)
	})

	Describe("FindCurrent", func() {
		Context("when a current vm cid is set", func() {
			BeforeEach(func() {
				err := repo.UpdateCurrent("fake-vm-cid")
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns current manifest sha1", func() {
				record, found, err := repo.FindCurrent()
				Expect(err).ToNot(HaveOccurred())
				Expect(found).To(BeTrue())
				Expect(record).To(Equal("fake-vm-cid"))
			})
		})

		Context("when a current vm cid is not set", func() {
			It("returns false", func() {
				_, found, err := repo.FindCurrent()
				Expect(err).ToNot(HaveOccurred())
				Expect(found).To(BeFalse())
			})
		})
	})

	Describe("UpdateCurrent", func() {
		It("updates vm cid", func() {
			err := repo.UpdateCurrent("fake-vm-cid")
			Expect(err).ToNot(HaveOccurred())

			deploymentConfig, err := configService.Load()
			Expect(err).ToNot(HaveOccurred())

			expectedConfig := DeploymentFile{
				CurrentVMCID: "fake-vm-cid",
			}
			Expect(deploymentConfig).To(Equal(expectedConfig))
		})
	})

	Describe("ClearCurrent", func() {
		It("updates vm cid", func() {
			err := repo.ClearCurrent()
			Expect(err).ToNot(HaveOccurred())

			deploymentConfig, err := configService.Load()
			Expect(err).ToNot(HaveOccurred())

			expectedConfig := DeploymentFile{
				CurrentVMCID: "",
			}
			Expect(deploymentConfig).To(Equal(expectedConfig))

			_, found, err := repo.FindCurrent()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeFalse())
		})
	})
})
