/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/zakkbob/mxguard/internal/config"

	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var cfgFile string
var Config config.Config

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "mxguard",
	Short:   "manage mxguard",
	Long:    `nothing here yet...`,
	Version: "0.1.0-alpha",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Hi! Nothing to see here yet...")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.WithError(err).Fatal("Failed to find home dir")
		}

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".mxguard")
		viper.SetConfigType("yaml")

	}

	viper.SetEnvPrefix("MXGUARD")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var err error
	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.WithError(err).Warning("Config file not found; continuing with defaults")
		} else {
			log.WithError(err).Fatal("Failed to read config file")
		}
	} else {
		log.WithFields(log.Fields{
			"config": viper.ConfigFileUsed(),
		}).Info("Found config file")
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.WithError(err).Fatal("Failed to unmarshal config into struct")
	}

	log.Info(Config)

	if Config.Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	for key, value := range viper.GetViper().AllSettings() {
		log.WithFields(log.Fields{
			key: value,
		}).Info("Using field")
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	RootCmd.PersistentFlags().BoolVarP(&Config.Verbose, "verbose", "v", false, "Display more verbose console output (default: false)")
	RootCmd.PersistentFlags().BoolVarP(&Config.Debug, "debug", "d", false, "Display debugging output in console (default: false)")
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
}
