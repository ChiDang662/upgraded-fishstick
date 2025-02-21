package main

import (
	"fmt"

	"upgraded-fishstick/config"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your application",
	Long:  `A longer description of your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, Cobra and Viper!")
	},
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.app.yaml)")
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
