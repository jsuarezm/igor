// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
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
	//"os/user"
    	"path/filepath"
	//"log"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	VERSION string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "igor",
	Short: "An infrastructure deployment assistance",
	Long: ` ▄█     ▄██████▄   ▄██████▄     ▄████████ 
███    ███    ███ ███    ███   ███    ███ 
███▌   ███    █▀  ███    ███   ███    ███ 
███▌  ▄███        ███    ███  ▄███▄▄▄▄██▀ 
███▌ ▀▀███ ████▄  ███    ███ ▀▀███▀▀▀▀▀   
███    ███    ███ ███    ███ ▀███████████ 
███    ███    ███ ███    ███   ███    ███ 
█▀     ████████▀   ▀██████▀    ███    ███ 
                               ███    ███ 
Greetings, this is IGOR, your assistance for infrastructure as a code
with terraform, packer and AWS.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	VERSION = version

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() { 
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.igor.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
        // Find home directory.
        home, err := homedir.Dir()
        if err != nil {
        	fmt.Println(err)
                os.Exit(1)
        }
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} 
	// Search config in home directory with name ".igor" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".igor")

	//viper.AutomaticEnv() // read in environment variables that match
	filename := filepath.Join(home, fmt.Sprintf("%s.yaml",".igor"))
	if err := viper.ReadInConfig(); err != nil {
    		_, err := os.Create(filename)
    		if err != nil {
        		panic(fmt.Sprintf("Failed to create %s", filename))
    		}
	}

	if err := viper.ReadInConfig(); err != nil {
    		panic(fmt.Sprintf("Created %s, but Viper failed to read it: %s",filename, err))
 	} 			

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
