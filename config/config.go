/* Copyright © 2018 Peter Alexander <peter.alexander@prodatalab.com>
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.*/

package config 

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// Config the app's configuration
type Config struct {
	Port int64 
	Config string 
	LogConfig LoggingConfig
}

// LoadConfig loads the config from a file if specified, otherwise
// from the environment
func LoadConfig(cmd *cobra.Command) (*Config, error) {
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil, err 
	}
	viper.SetEnvPrefix("CODE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/.codectl")
	}
	if err := viper.ReadInConfig(); err != nil {
		return nil, err 
	}
	return populateConfig(new(Config))

}