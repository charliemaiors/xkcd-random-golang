// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/charliemaiors/xkcd-random-golang/xkcd"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "xkcd-random-golang",
	Short: "This is a simple webserver for xkcd random comic retrieving ",
	Long: `This web server uses letsencrypt for random comic retrieve, is defined only for education purposes in order to 
	show how letsencrypt and acme works`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		xkcd.RunSrv()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(viper.AutomaticEnv)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().String("domain", "example.com", "Specify your dns name where application is currently listening")
	viper.BindPFlag("domain", RootCmd.PersistentFlags().Lookup("domain"))
	viper.SetDefault("domain", "example.com")

	dir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	RootCmd.PersistentFlags().String("certdir", dir, "Default directory for certificates")
	viper.BindPFlag("certdir", RootCmd.PersistentFlags().Lookup("certdir"))
	viper.SetDefault("certdir", dir)
}
