package mod

import (
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

func getDepStr(line string) (string, string, bool) {
	// github.com/hqpko/gomod github.com/urfave/cli@v1.20.0
	ss := strings.Split(line, " ")
	if len(ss) != 2 {
		return "", "", false
	}
	return ss[0], ss[1], true
}

func getDeep(c *cli.Context) int {
	deep, _ := strconv.Atoi(c.Args().First())
	if deep <= 0 {
		deep = 999
	}
	return deep
}

func isRoot(path string) bool {
	return strings.Index(path, "@") == -1
}
