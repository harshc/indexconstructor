package cmd

import (
	"os"
	"github.com/spf13/cobra"
	log "github.com/Sirupsen/logrus"
)

const configDirName = ".indexconstructor"
var cfgFile string
var version string

var rootCmd = &cobra.Command{
	Use: "indexconstructor",
	Short: "Constructs the dataset for the index processor",
	Long: "Constructs a JSON dataset for the custom index processor",
	Run: func(cmd *cobra.Command, args []string) {
		// nothing do here - no default action
	},
}

// Execute adds all child commands to the root command
// sets flags appropriately. This is called by the main().
// It only needs to happen once.
func Execute(ver string){
	version = ver
	if err := rootCmd.Execute(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatalln("Error executing root command")
		os.Exit(-1)
	}
}

func init() {
	// define flags
}