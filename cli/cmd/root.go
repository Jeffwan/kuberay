/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/ray-project/ray-contrib/cli/pkg/cmd/info"
	"github.com/ray-project/ray-contrib/cli/pkg/cmd/cluster"
	"github.com/ray-project/ray-contrib/cli/pkg/cmd/runtime"
	"github.com/ray-project/ray-contrib/cli/pkg/cmd/version"
	"os"

	"github.com/spf13/cobra"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/kris-nova/logger"
	lol "github.com/kris-nova/lolgopher"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	loggerLevel := rootCmd.PersistentFlags().IntP("verbose", "v", 3, "set log level, use 0 to silence, 4 for debugging and 5 for debugging with AWS debug logging")
	colorValue := rootCmd.PersistentFlags().StringP("color", "C", "true", "toggle colorized logs (valid options: true, false, fabulous)")
	cobra.OnInitialize(initConfig, func() {
		initLogger(*loggerLevel, *colorValue)
	})

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().BoolP("help", "h", false, "help for this command")
	rootCmd.AddCommand(info.NewCmdInfo())
	rootCmd.AddCommand(version.NewCmdVersion())
	rootCmd.AddCommand(cluster.NewCmdCluster())
	rootCmd.AddCommand(runtime.NewCmdRuntime())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func initLogger(level int, colorValue string) {
	logger.Layout = "2021-01-02 15:04:05"

	var bitwiseLevel int
	switch level {
	case 4:
		bitwiseLevel = logger.LogDeprecated | logger.LogAlways | logger.LogSuccess | logger.LogCritical | logger.LogWarning | logger.LogInfo | logger.LogDebug
	case 3:
		bitwiseLevel = logger.LogDeprecated | logger.LogAlways | logger.LogSuccess | logger.LogCritical | logger.LogWarning | logger.LogInfo
	case 2:
		bitwiseLevel = logger.LogDeprecated | logger.LogAlways | logger.LogSuccess | logger.LogCritical | logger.LogWarning
	case 1:
		bitwiseLevel = logger.LogDeprecated | logger.LogAlways | logger.LogSuccess | logger.LogCritical
	case 0:
		bitwiseLevel = logger.LogDeprecated | logger.LogAlways | logger.LogSuccess
	default:
		bitwiseLevel = logger.LogDeprecated | logger.LogEverything
	}
	logger.BitwiseLevel = bitwiseLevel

	switch colorValue {
	case "fabulous":
		logger.Writer = lol.NewLolWriter()
	case "true":
		logger.Writer = color.Output
	}

	logger.Line = func(prefix, format string, a ...interface{}) string {
		if !strings.Contains(format, "\n") {
			format = fmt.Sprintf("%s%s", format, "\n")
		}
		now := time.Now()
		fNow := now.Format(logger.Layout)
		var colorize func(format string, a ...interface{}) string
		var icon string
		switch prefix {
		case logger.PreAlways:
			icon = "✿"
			colorize = color.GreenString
		case logger.PreCritical:
			icon = "✖"
			colorize = color.RedString
		case logger.PreInfo:
			icon = "ℹ"
			colorize = color.CyanString
		case logger.PreDebug:
			icon = "▶"
			colorize = color.GreenString
		case logger.PreSuccess:
			icon = "✔"
			colorize = color.CyanString
		case logger.PreWarning:
			icon = "!"
			colorize = color.GreenString
		default:
			icon = "ℹ"
			colorize = color.CyanString
		}

		out := fmt.Sprintf(format, a...)
		out = fmt.Sprintf("%s [%s]  %s", fNow, icon, out)
		if colorValue == "true" {
			out = colorize(out)
		}

		return out
	}
}
