package cmd

import (
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "github.com/Sirupsen/logrus"
)

var cfgFile string
const configDirName = ".indexconstructor"
var rootCmd = &cobra.Command{
	Use: "construct",
	Short: "Constructs the dataset for the index processor",
	Long: "Constructs a JSON dataset for the custom index processor",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cobra.OnInitialize(initConfig)
	// define flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.indexconstructor/config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config") // name of the config file (without extension)
		viper.AddConfigPath(filepath.Join(os.TempDir(), configDirName))
		viper.AddConfigPath(filepath.Join("$HOME", configDirName))
		viper.AddConfigPath(".")
		viper.AutomaticEnv() // read in any other variables that match
	}

	// if we find a config file, read it in
	if err := viper.ReadInConfig(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"filepath": cfgFile,
		}).Fatalln("Error reading config file")
	}
}