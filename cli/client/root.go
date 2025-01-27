package client

import (
	"fmt"
	"goSum/cli/client/sumCmd"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// rootCmd represents the base command when cli is called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gosum",
	Short: "Perform mathematical calculations from the command line",
	Run: func(cmd *cobra.Command, args []string) {
		// Print logo
		whiteBold := color.New(color.FgHiWhite, color.Bold)
		logo, err := ioutil.ReadFile("logo.txt")
		if err != nil {
			log.Fatalln("Error opening 'logo.txt':", err)
		}
		whiteBold.Println(string(logo))
		fmt.Println("Perform mathematical operations from the command line!")
		fmt.Print("E.g. 'sum -n 1,3,4,6'\n")
	},
}

// Initialize cobra and add subcommand 'sum'
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.omniactl.yaml)")

	// Add more subcommands for different calculations here e.g. minusCmd or moduloCmd
	sumCmd.AddSubCommands(rootCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".gosum")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
