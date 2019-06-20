package mod

import (
	"os"

	"github.com/hqpko/gosh"
	"github.com/urfave/cli"
)

const (
	defProxyURL = "https://athens.azurefd.net"
)

func ActionGoModTidy(c *cli.Context) error {
	if e := initEnvProxy(); e != nil {
		return e
	}
	return gosh.Run("go mod tidy")
}

func initEnvProxy() error {
	if os.Getenv("GOPROXY") == "" {
		return os.Setenv("GOPROXY", defProxyURL)
	}
	return nil
}
