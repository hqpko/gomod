package mod

import (
	"strings"

	"github.com/hqpko/gosh"
	"github.com/urfave/cli"
)

func ActionOrigin(c *cli.Context) error {
	cmds := append([]string{"go", "mod", c.Command.Name}, c.Args()...)
	return gosh.Run(strings.Join(cmds, " "))
}
