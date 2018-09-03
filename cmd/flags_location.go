/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
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

package cmd

import (
	"github.com/mysteriumnetwork/node/core/node"
	"github.com/urfave/cli"
)

var (
	ipifyUrlFlag = cli.StringFlag{
		Name:  "ipify-url",
		Usage: "Address (URL form) of ipify service",
		Value: "https://api.ipify.org/",
	}
	locationDatabaseFlag = cli.StringFlag{
		Name:  "location.database",
		Usage: "Service location autodetect database of GeoLite2 format e.g. http://dev.maxmind.com/geoip/geoip2/geolite2/",
		Value: "GeoLite2-Country.mmdb",
	}
	locationCountryFlag = cli.StringFlag{
		Name:  "location.country",
		Usage: "Service location country. If not given country is autodetected",
		Value: "",
	}
)

// RegisterLocationFlags function register location flags to flag list
func RegisterLocationFlags(flags *[]cli.Flag) {
	*flags = append(*flags, ipifyUrlFlag, locationDatabaseFlag, locationCountryFlag)
}

// ParseLocationFlags function fills in location options from CLI context
func ParseLocationFlags(ctx *cli.Context) node.LocationOptions {
	return node.LocationOptions{
		ctx.GlobalString(ipifyUrlFlag.Name),
		ctx.GlobalString(locationDatabaseFlag.Name),
		ctx.GlobalString(locationCountryFlag.Name),
	}
}
