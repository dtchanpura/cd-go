package config

import (
	"fmt"
	"os"

	"github.com/dtchanpura/deployment-agent/constants"
)

// StoredProjects Contains the latest stored projects
var StoredProjects []Project

// StoredServe Contains the latest stored serve details
var StoredServe Serve

// var StoredConfiguration Configuration

// InitializeConfiguration for initializing the configuration file.
func InitializeConfiguration(cfgFile string, overwrite bool) {
	defaultConfig := Configuration{
		ServeConfig: Serve{
			Host: "",
			Port: 8000,
		},
	}

	err := UpdateConfiguration(cfgFile, defaultConfig, overwrite)
	if err != nil {
		if err.Error() == constants.ErrorFileExists {
			fmt.Println("Configuration already initialized. Use -f for overwriting it forcefully.")
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Printf("Initializing file with defaults: %s\n", cfgFile)
}
