/* Copyright © 2018 Peter Alexander <peter.alexander@prodatalab.com>
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.*/

package config 

import (
	"bufio"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
)

// LoggingConfig specifies all the parameters needed for logging
type LoggingConfig struct {
	Level string 
	File string 
}

// ConfigureLogging will take the logging configuration and also adds
// a few default parameters
func (lc *LoggingConfig) ConfigureLogging() (*logrus.Entry, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err 
	}
	if config.File != "" {
		f, errOpen := os.OpenFile(config.File, os.O_RDWR|os.O_APPEND, 0660)
		if errOpen != nil {
			return nil, errOpen
		}
		logrus.SetOutput(bufio.NewWriter(f))
	}
	if config.Level != "" {
		level, err := logrus.ParseLevel(strings.ToUpper((config.Level)))
		if err != nil {
			return nil, err 
		}
		logrus.SetLevel(level)
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableTimestamp: false,
	})
	return logrus.StandardLogger().WithField("hostname", hostname), nil
}