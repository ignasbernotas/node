/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mysteriumnetwork/node/cmd"
	command_cli "github.com/mysteriumnetwork/node/cmd/commands/cli"
	"github.com/mysteriumnetwork/node/cmd/commands/license"
	"github.com/mysteriumnetwork/node/cmd/commands/run"
	"github.com/mysteriumnetwork/node/cmd/commands/service"
	"github.com/mysteriumnetwork/node/cmd/commands/version"
	"github.com/mysteriumnetwork/node/core/node"
	"github.com/mysteriumnetwork/node/metadata"
	tequilapi_client "github.com/mysteriumnetwork/node/tequilapi/client"
	"github.com/mysteriumnetwork/node/utils"
	"github.com/urfave/cli"
)

func main() {
	app, err := NewCommand()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// NewCommand function creates application master command
func NewCommand() (*cli.App, error) {
	cli.VersionPrinter = func(ctx *cli.Context) {
		versionCommand.Run(ctx)
	}

	app := cli.NewApp()
	app.Usage = "VPN server and client for Mysterium Network https://mysterium.network/"
	app.Authors = []cli.Author{
		{`The "MysteriumNetwork/node" Authors`, "mysterium-dev@mysterium.network"},
	}
	app.Version = metadata.VersionAsString()
	app.Copyright = licenseCopyright
	if err := cmd.RegisterFlagsNode(&app.Flags); err != nil {
		return nil, err
	}
	app.Flags = append(app.Flags, cliFlag)
	app.Commands = []cli.Command{
		*versionCommand,
		*license.NewCommand(licenseCopyright),
		*service.NewCommand(),
	}
	app.Action = runMain

	return app, nil
}

func runMain(ctx *cli.Context) error {
	options := commandOptions{
		CLI:         ctx.Bool(cliFlag.Name),
		NodeOptions: cmd.ParseFlagsNode(ctx),
	}

	if options.CLI {
		return runCLI(options.NodeOptions)
	}

	fmt.Println(versionSummary)
	fmt.Println()

	return run.NewCommand().Run(ctx)
}

func runCLI(options node.Options) error {
	cmdCli := command_cli.NewCommand(
		filepath.Join(options.Directories.Data, ".cli_history"),
		tequilapi_client.NewClient(options.TequilapiAddress, options.TequilapiPort),
	)
	cmd.RegisterSignalCallback(utils.HardKiller(cmdCli.Kill))

	return cmdCli.Run()
}

type commandOptions struct {
	CLI         bool
	NodeOptions node.Options
}

var (
	licenseCopyright = metadata.LicenseCopyright(
		"run command 'license --warranty'",
		"run command 'license --conditions'",
	)
	versionSummary = metadata.VersionAsSummary(licenseCopyright)
	versionCommand = version.NewCommand(versionSummary)

	cliFlag = cli.BoolFlag{
		Name:  "cli",
		Usage: "Run an interactive CLI based Mysterium UI",
	}
)
