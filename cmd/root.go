/*
Copyright © 2022 Joakim Karlsson

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
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wifiqr",
	Short: "Generate a Wi-Fi QR code",
	RunE: func(cmd *cobra.Command, args []string) error {

		b, err := genWiFiQR(getFlags(cmd))
		if err != nil {
			return err
		}

		if err := os.WriteFile("qr.png", b, 0755); err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	pf := rootCmd.PersistentFlags()
	pf.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	lf := rootCmd.Flags()
	lf.BoolP("toggle", "t", false, "Help message for toggle")
	lf.StringP("authenticationType", "a", "WPA", "AuthenticationType")
	lf.StringP("ssid", "i", "Test", "SSID")
	lf.StringP("password", "p", "123456", "Password")
	lf.BoolP("hidden", "x", false, "Hidden SSID")
	lf.IntP("size", "s", 1000, "Set QR size")
}

func getFlags(cmd *cobra.Command) (string, string, string, bool, int) {
	lf := cmd.LocalFlags()

	typ, err := lf.GetString("authenticationType")
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	ssid, err := lf.GetString("ssid")
	if err != nil {
		log.Fatal(err)
	}

	passwd, err := lf.GetString("password")
	if err != nil {
		log.Fatal(err)
	}

	hidden, err := lf.GetBool("hidden")
	if err != nil {
		log.Fatal(err)
	}

	size, err := lf.GetInt("size")
	if err != nil {
		log.Fatal(err)
	}
	return typ, ssid, passwd, hidden, size
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

		// Search config in home directory with name ".cmd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
