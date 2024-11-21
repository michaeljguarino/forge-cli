package up

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pluralsh/plural-cli/pkg/utils"
	"github.com/pluralsh/plural-cli/pkg/utils/git"
)

type templatePair struct {
	from      string
	to        string
	overwrite bool
	cloud     bool
	cloudless bool
}

func (ctx *Context) Generate() (dir string, err error) {
	dir, err = os.MkdirTemp("", "sampledir")
	ctx.dir = dir
	if err != nil {
		return
	}

	if err = git.PathClone("https://github.com/pluralsh/bootstrap.git", "template-branch", dir); err != nil {
		return
	}

	prov := ctx.Provider.Name()
	tpls := []templatePair{
		{from: ctx.path("charts/runtime/values.yaml.tpl"), to: "./helm-values/runtime.yaml", overwrite: true},
		{from: ctx.path("helm/certmanager.yaml"), to: "./helm-values/certmanager.yaml", overwrite: true},
		{from: ctx.path("helm/flux.yaml"), to: "./helm-values/flux.yaml", overwrite: true},
		{from: ctx.path(fmt.Sprintf("templates/providers/bootstrap/%s.tf", prov)), to: "terraform/mgmt/provider.tf"},
		{from: ctx.path(fmt.Sprintf("templates/setup/providers/%s.tf", prov)), to: "terraform/mgmt/mgmt.tf"},
		{from: ctx.path("templates/setup/console.tf"), to: "terraform/mgmt/console.tf", cloudless: true},
		{from: ctx.path(fmt.Sprintf("templates/providers/apps/%s.tf", prov)), to: "terraform/apps/provider.tf", cloudless: true},
		{from: ctx.path("templates/providers/apps/cloud.tf"), to: "terraform/apps/provider.tf", cloud: true},
		{from: ctx.path("templates/setup/cd.tf"), to: "terraform/apps/cd.tf"},
		{from: ctx.path("README.md"), to: "README.md", overwrite: true},
	}

	for _, tpl := range tpls {
		if utils.Exists(tpl.to) && !tpl.overwrite {
			fmt.Printf("%s already exists, skipping for now...\n", tpl.to)
			continue
		}

		if tpl.cloudless && ctx.Cloud {
			continue
		}

		if tpl.cloud && !ctx.Cloud {
			continue
		}

		if err = ctx.templateFrom(tpl.from, tpl.to); err != nil {
			return
		}
	}

	copies := []templatePair{
		{from: ctx.path("terraform/modules/clusters"), to: "terraform/modules/clusters"},
		{from: ctx.path(fmt.Sprintf("terraform/clouds/%s", prov)), to: "terraform/mgmt/cluster"},
		{from: ctx.path("setup"), to: "bootstrap"},
		{from: ctx.path("templates"), to: "templates"},
		{from: ctx.path("resources"), to: "resources"},
	}

	for _, copy := range copies {
		if utils.Exists(copy.to) && !copy.overwrite {
			continue
		}

		if err = utils.CopyDir(copy.from, copy.to); err != nil {
			return
		}
	}

	if ctx.Cloud {
		toRemove := []string{"bootstrap/console.yaml", "bootstrap/flux.yaml"}
		for _, f := range toRemove {
			os.Remove(f)
		}
	}

	ctx.changeDelims()
	overwrites := []templatePair{
		{from: "resources/monitoring/services", to: "resources/monitoring/services"},
		{from: "resources/policy/services", to: "resources/policy/services"},
		{from: "bootstrap", to: "bootstrap"},
	}

	for _, tpl := range overwrites {
		if utils.IsDir(tpl.from) {
			files, err := utils.ListDirectory(tpl.from)
			if err != nil {
				return dir, err
			}

			for _, file := range files {
				destFile, err := filepath.Rel(tpl.from, file)
				if err != nil {
					return dir, err
				}

				destFile = filepath.Join(string(tpl.to), destFile)
				if err = ctx.templateFrom(file, destFile); err != nil {
					return dir, err
				}
			}

			continue
		}

		if err = ctx.templateFrom(tpl.from, tpl.to); err != nil {
			return
		}
	}

	return
}

func (ctx *Context) afterSetup() error {
	prov := ctx.Provider.Name()
	overwrites := []templatePair{
		{from: ctx.path(fmt.Sprintf("templates/setup/stacks/%s.yaml", prov)), to: "bootstrap/stacks/serviceaccount.yaml"},
	}

	ctx.Delims = nil
	for _, tpl := range overwrites {
		if err := ctx.templateFrom(tpl.from, tpl.to); err != nil {
			return err
		}
	}

	return nil
}

func (ctx *Context) path(p string) string {
	return filepath.Join(ctx.dir, p)
}
