/* Copyright © 2018 Peter Alexander <peter.alexander@prodatalab.com>
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.*/

package cmd

import (
	"fmt"

	"github.com/ProDataLab/cbp"
	"github.com/spf13/cobra"
)

// componentCmd represents the component command
var (
	name      string
	url       string
	port      string
	socket    string
	transport string

	componentCmd = &cobra.Command{
		Use:   "component",
		Short: "compoents are the heart of codedepot",
		Long: `components are the heart of codedepot
		
		Find more information at https://codedepot.tech/docs/component.

		TODO: add a nice lengthy overview here
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("component called. %v\n", args)
			c, err := cbp.NewComponent("create-component")
			if err != nil {
				panic(err)
			}
			c.AddSocket("")
		},
	}
)

func init() {
	createCmd.AddCommand(componentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// componentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// componentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	componentCmd.Flags().StringVarP(&socket, "socket", "s", "req", "Socket Type: req|rep|push|pull|pub|sub")
}
