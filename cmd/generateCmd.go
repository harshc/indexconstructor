package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/harshc/indexconstructor/globals"
	"math/rand"
	"os"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"encoding/json"
)

var outFile string

var generateCmd = &cobra.Command{
	Use: "generate",
	Short: "Generates a json file from a source word list",
	Long: "Reads dummy word list from the included yaml file and outputs a json file.",
	Run: func(cmd *cobra.Command, args []string) {
		var allWords []string
		err := viper.UnmarshalKey("words", &allWords)
		log.WithFields(log.Fields{
			"error": err,
		}).Errorln("Error unmarshalling the mock data yaml")

		jsonFile, _ := os.Create(outFile)

		defer jsonFile.Close()
		log.WithFields(log.Fields{
			"filename": outFile,
		}).Debugln("Writing to file")

		for _, w := range allWords {
			word := &globals.ScoredName{
				Name: w,
				Score: rand.Intn(1000),
			}

			jsonString,_ := json.Marshal(word)
			fmt.Fprintln(jsonFile, string(jsonString))
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&outFile, "outfile", "o", "dataset.json", "name of the file to write to")
}
