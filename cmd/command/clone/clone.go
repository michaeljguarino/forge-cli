package clone

import (
	"github.com/pluralsh/plural-cli/pkg/common"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:      "clone",
		Usage:     "clones and decrypts a plural repo",
		ArgsUsage: "URL",
		Action:    common.HandleClone,
	}
}
