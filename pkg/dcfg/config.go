/*
 * File: config.go
 * Date: Wednesday, 14th August 2024 10:21:50 PM
 * Author: Hy Nguyen
 * Copyright - 2024, HYSOFTWARE LLC
 */

package dcfg

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	local = "local"
)

var (
	rtEnvvar         = "RUNTIME_ENVVAR"
	settingsFileName = "settings"

	Cfg     *viper.Viper
	verbose bool = true
)

// SetRuntimeEnv sets the runtime environment envar
// This envar is used to get the runtime environment,
// e.g prod or staging or devint
func SetRuntimeEnv(newvar string) {
	rtEnvvar = newvar
}

// SetVerbose sets the flag to enable/disable the verbosity
func SetVerbose(f bool) {
	verbose = f
}

// InitConfig initializes the config module using the
// given cfgDir to load config files
func InitConfig(cfgDir string) error {

	if verbose {
		fmt.Println("----------------------------------")
		fmt.Println("Initialize dcfg")
	}
	Cfg = viper.New()
	Cfg.AutomaticEnv()
	Cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if cfgDir == "" {
		cfgDir = "."
	}

	if verbose {
		fmt.Printf("...searching for config file in directory=%s\n", cfgDir)
	}

	Cfg.AddConfigPath(cfgDir)
	Cfg.SetConfigType("yaml")

	files := getFileNames(rtEnvvar)

	for _, f := range files {
		Cfg.SetConfigName(f)
		if verbose {
			fmt.Printf("...loading config file %s.yaml\n", f)
		}
		err := Cfg.MergeInConfig()
		if err != nil {
			_, isFnfErr := err.(viper.ConfigFileNotFoundError)
			if isFnfErr {
				fmt.Printf("	file not found - %s.yanl\n", f)
			}
		}
	}
	return nil
}

func getFileNames(rtEnvvar string) []string {
	env := strings.ToLower((Cfg.GetString(rtEnvvar)))

	if env == "" {
		env = local
	}
	return []string{settingsFileName, fmt.Sprintf("%s_%s", settingsFileName, env)}
}
