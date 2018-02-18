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
var allWords []string

// this is a helper method to create a mock dataset
var generateCmd = &cobra.Command{
	Use: "generate",
	Short: "Generates a json file from a source word list",
	Long: "Reads dummy word list from the included yaml file and outputs a json file.",
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.UnmarshalKey("words", &allWords)
		log.WithFields(log.Fields{
			"error": err,
		}).Errorln("Error unmarshalling the mock data yaml")
		go CreateJsonFile(allWords, outFile)
	},
}

// CreateJsonFile creates a json file with input word list
func CreateJsonFile(allWords []string, outfile string) {
	jsonFile, _ := os.Create(outfile)

	defer jsonFile.Close()
	log.WithFields(log.Fields{
		"filename": outFile,
	}).Debugln("Writing to file")

	var words []globals.ScoredName
	for _, w := range allWords {
		word := &globals.ScoredName{
			Name: w,
			Score: rand.Intn(1000),
		}

		words = append(words, *word)
	}

	jsonString,_ := json.Marshal(words)
	fmt.Fprintln(jsonFile, string(jsonString))
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&outFile, "outfile", "o", "dataset.json", "name of the file to write to")
}
