package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"strings"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version of indexconstructor",
	Long: "All software has versions, this is mine",
	Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(strings.Join([]string{rootCmd.Use, version}, ":"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
