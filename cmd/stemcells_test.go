package cmd_test

import (
	"errors"

	semver "github.com/cppforlife/go-semi-semantic/version"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-cli/cmd"
	boshdir "github.com/cloudfoundry/bosh-cli/director"
	fakedir "github.com/cloudfoundry/bosh-cli/director/fakes"
	fakeui "github.com/cloudfoundry/bosh-cli/ui/fakes"
	boshtbl "github.com/cloudfoundry/bosh-cli/ui/table"
)

var _ = Describe("StemcellsCmd", func() {
	var (
		ui       *fakeui.FakeUI
		director *fakedir.FakeDirector
		command  StemcellsCmd
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		director = &fakedir.FakeDirector{}
		command = NewStemcellsCmd(ui, director)
	})

	Describe("Run", func() {
		act := func() error { return command.Run() }

		It("lists stemcells", func() {
			stemcells := []boshdir.Stemcell{
				&fakedir.FakeStemcell{
					NameStub:        func() string { return "stem1" },
					VersionStub:     func() semver.Version { return semver.MustNewVersionFromString("stem1-ver") },
					VersionMarkStub: func(mark string) string { return "stem1-ver-mark" + mark },
					OSNameStub:      func() string { return "stem1-os" },
					CIDStub:         func() string { return "stem1-cid" },
				},
				&fakedir.FakeStemcell{
					NameStub:    func() string { return "stem2" },
					VersionStub: func() semver.Version { return semver.MustNewVersionFromString("stem2-ver") },
					OSNameStub:  func() string { return "stem2-os" },
					CIDStub:     func() string { return "stem2-cid" },
				},
				&fakedir.FakeStemcell{
					NameStub:    func() string { return "stem3" },
					VersionStub: func() semver.Version { return semver.MustNewVersionFromString("stem3-ver") },
					OSNameStub:  func() string { return "stem3-os" },
					CIDStub:     func() string { return "stem3-cid" },
				},
				&fakedir.FakeStemcell{
					NameStub:    func() string { return "stem4" },
					VersionStub: func() semver.Version { return semver.MustNewVersionFromString("stem4-ver") },
					CIDStub:     func() string { return "stem4-cid" },
				},
			}

			director.StemcellsReturns(stemcells, nil)

			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(ui.Table).To(Equal(boshtbl.Table{
				Content: "stemcells",

				Header: []string{"Name", "Version", "OS", "CID"},

				SortBy: []boshtbl.ColumnSort{
					{Column: 0, Asc: true},
					{Column: 1, Asc: false},
				},

				Rows: [][]boshtbl.Value{
					{
						boshtbl.NewValueString("stem1"),
						boshtbl.NewValueSuffix(
							boshtbl.NewValueVersion(semver.MustNewVersionFromString("stem1-ver")),
							"stem1-ver-mark*",
						),
						boshtbl.NewValueString("stem1-os"),
						boshtbl.NewValueString("stem1-cid"),
					},
					{
						boshtbl.NewValueString("stem2"),
						boshtbl.NewValueSuffix(
							boshtbl.NewValueVersion(semver.MustNewVersionFromString("stem2-ver")),
							"",
						),
						boshtbl.NewValueString("stem2-os"),
						boshtbl.NewValueString("stem2-cid"),
					},
					{
						boshtbl.NewValueString("stem3"),
						boshtbl.NewValueSuffix(
							boshtbl.NewValueVersion(semver.MustNewVersionFromString("stem3-ver")),
							"",
						),
						boshtbl.NewValueString("stem3-os"),
						boshtbl.NewValueString("stem3-cid"),
					},
					{
						boshtbl.NewValueString("stem4"),
						boshtbl.NewValueSuffix(
							boshtbl.NewValueVersion(semver.MustNewVersionFromString("stem4-ver")),
							"",
						),
						boshtbl.NewValueString(""),
						boshtbl.NewValueString("stem4-cid"),
					},
				},

				Notes: []string{"(*) Currently deployed"},
			}))
		})

		It("returns error if stemcells cannot be retrieved", func() {
			director.StemcellsReturns(nil, errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
