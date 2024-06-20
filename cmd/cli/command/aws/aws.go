package aws

import (
	"github.com/common-fate/cli/cmd/cli/command/aws/rdsv2"
	"github.com/urfave/cli/v2"
)

var Command = cli.Command{
	Name:  "aws",
	Usage: "Perform AWS Operations",
	Subcommands: []*cli.Command{
		&rdsv2.Command,
	},
}
