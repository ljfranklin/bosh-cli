package cmd

import (
	cmdconf "github.com/cloudfoundry/bosh-init/cmd/config"
	boshui "github.com/cloudfoundry/bosh-init/ui"
	boshtbl "github.com/cloudfoundry/bosh-init/ui/table"
)

type EnvironmentsCmd struct {
	config cmdconf.Config
	ui     boshui.UI
}

func NewEnvironmentsCmd(config cmdconf.Config, ui boshui.UI) EnvironmentsCmd {
	return EnvironmentsCmd{config: config, ui: ui}
}

func (c EnvironmentsCmd) Run() error {
	environments := c.config.Environments()

	table := boshtbl.Table{
		Content: "targets",
		Header:  []string{"URL", "Alias"},
		SortBy:  []boshtbl.ColumnSort{{Column: 0, Asc: true}},
	}

	for _, t := range environments {
		table.Rows = append(table.Rows, []boshtbl.Value{
			boshtbl.NewValueString(t.URL),
			boshtbl.NewValueString(t.Alias),
		})
	}

	c.ui.PrintTable(table)

	return nil
}