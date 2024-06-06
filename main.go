package main

import (
	"dev-docs-cli/cmd"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Get the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home directory, %s", err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(filepath.Join(homeDir, ".devdocs"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	cmd.Execute()
}

//package main
//
//import (
//	"github.com/mwiater/golangcliscaffold/cmd"
//	"github.com/spf13/viper"
//	"log"
//)
//
//func main() {
//	viper.SetConfigName("config")
//	viper.SetConfigType("yaml")
//	viper.AddConfigPath(".")
//	viper.AddConfigPath("$HOME/.devdocs")
//
//	if err := viper.ReadInConfig(); err != nil {
//		log.Fatalf("Error reading config file, %s", err)
//	}
//
//	cmd.Execute()
//}
