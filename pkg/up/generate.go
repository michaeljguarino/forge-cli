package up

import (
	"fmt"
	"os"

	"github.com/pluralsh/plural-cli/pkg/utils"
	"github.com/pluralsh/plural-cli/pkg/utils/git"
)

type templatePair struct {
	from      string
	to        string
	overwrite bool
}

func (ctx *Context) Cleanup() {
	git.RemoveSubmodule("bootstrap")
	os.RemoveAll("./bootstrap")
}

func (ctx *Context) Generate() error {
	if !utils.Exists("./bootstrap") {
		if err := git.BranchedSubmodule("https://github.com/pluralsh/bootstrap.git", "stacks-support"); err != nil {
			return err
		}
	}

	prov := ctx.Provider.Name()
	tpls := []templatePair{
		{from: "./bootstrap/charts/runtime/values.yaml.tpl", to: "./helm-values/runtime.yaml", overwrite: true},
		{from: "./bootstrap/helm/certmanager.yaml", to: "./helm-values/certmanager.yaml", overwrite: true},
		{from: "./bootstrap/helm/flux.yaml", to: "./helm-values/flux.yaml", overwrite: true},
		{from: fmt.Sprintf("./bootstrap/templates/providers/bootstrap/%s.tf", prov), to: "clusters/provider.tf"},
		{from: fmt.Sprintf("./bootstrap/templates/setup/providers/%s.tf", prov), to: "clusters/mgmt.tf"},
		{from: "./bootstrap/templates/setup/console.tf", to: "clusters/console.tf"},
		{from: fmt.Sprintf("./bootstrap/templates/providers/apps/%s.tf", prov), to: "apps/terraform/provider.tf"},
		{from: "./bootstrap/templates/setup/cd.tf", to: "apps/terraform/cd.tf"},
		{from: "./bootstrap/README.md", to: "README.md", overwrite: true},
	}

	for _, tpl := range tpls {
		if utils.Exists(tpl.to) && !tpl.overwrite {
			fmt.Printf("%s already exists, skipping for now...\n", tpl.to)
			continue
		}

		if err := ctx.templateFrom(tpl.from, tpl.to); err != nil {
			return err
		}
	}

	copies := []templatePair{
		{from: "./bootstrap/terraform/modules/clusters", to: "terraform/modules/clusters"},
		{from: fmt.Sprintf("./bootstrap/terraform/clouds/%s", prov), to: "terraform/modules/mgmt"},
		{from: "./bootstrap/apps/repositories", to: "apps/repositories"},
		{from: "./bootstrap/apps/services", to: "apps/services"},
		{from: "./bootstrap/templates", to: "templates"},
	}

	for _, copy := range copies {
		if utils.Exists(copy.to) && !copy.overwrite {
			continue
		}

		if err := utils.CopyDir(copy.from, copy.to); err != nil {
			return err
		}
	}

	ctx.changeDelims()
	overwrites := []templatePair{
		{from: "apps/services/setup.yaml", to: "apps/services/setup.yaml"},
		{from: "apps/services/pr-automation/cluster-creator.yaml", to: "apps/services/pr-automation/cluster-creator.yaml"},
	}

	for _, tpl := range overwrites {
		if err := ctx.templateFrom(tpl.from, tpl.to); err != nil {
			return err
		}
	}

	return nil
}

func (ctx *Context) afterSetup() error {
	prov := ctx.Provider.Name()
	overwrites := []templatePair{
		{from: fmt.Sprintf("./bootstrap/templates/setup/stacks/%s.yaml", prov), to: "apps/services/stacks/serviceaccount.yaml"},
	}

	ctx.Delims = nil
	for _, tpl := range overwrites {
		if err := ctx.templateFrom(tpl.from, tpl.to); err != nil {
			return err
		}
	}

	return nil
}
