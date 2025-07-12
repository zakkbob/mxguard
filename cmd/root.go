/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	"github.com/zakkbob/mxguard/internal/config"

	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var Logger zerolog.Logger
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
		Logger.Info().Msg("Hi! Nothing to see here yet...")
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
	if Config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			Logger.Fatal().Err(err).Msg("Failed to find home dir")
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
			Logger.Warn().Err(err).Msg("Config file not found; continuing with defaults")
		} else {
			Logger.Fatal().Err(err).Msg("Failed to read config file")
		}
	} else {
		Logger.Debug().Str("config", viper.ConfigFileUsed()).Msg("Found config file")
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		Logger.Fatal().Err(err).Msg("Failed to unmarshal config into struct")
	}

	Logger.Debug().Any("config", Config).Msg("Config struct")

	if Config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	for key, value := range viper.GetViper().AllSettings() {
		Logger.Debug().Any(key, value).Msg("Using field")
	}
}

func initLogger() {
	buildInfo, _ := debug.ReadBuildInfo()

	Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s |", i)
		},
		FormatCaller: func(i interface{}) string {
			return filepath.Base(fmt.Sprintf("%s", i))
		},
		PartsExclude: []string{
			zerolog.TimestampFieldName,
		}}).Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()
//		Sample(&zerolog.BasicSampler{N: 5})
}

func init() {
	initLogger()
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	RootCmd.PersistentFlags().BoolVarP(&Config.Verbose, "verbose", "v", false, "Display more verbose console output (default: false)")
	RootCmd.PersistentFlags().BoolVarP(&Config.Debug, "debug", "d", false, "Display debugging output in console (default: false)")
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))
}
