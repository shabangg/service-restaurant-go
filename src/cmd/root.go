package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string

	// root command RootCmd
	RootCmd = &cobra.Command{
		Use:   "restaurant [OPTIONS] [COMMANDS]",
		Short: "Restaurant Service ",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is config.yaml)")
	RootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "rohanluthra <rohanluthra.1123@gmail.com>")
}

// Execute executes the root command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
