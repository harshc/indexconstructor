package cmd

import (
	"github.com/spf13/cobra"
	"github.com/harshc/indexconstructor/globals"
	"math/rand"
	"os"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"encoding/json"
	"bufio"
	"io"
	"strings"
)

var outFile string
var inFile string
// this is a helper method to create a mock dataset
var generateCmd = &cobra.Command{
	Use: "generate",
	Short: "Generates a json file from a source word list",
	Long: "Reads dummy word list from the included yaml file and outputs a json file.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infoln(outFile)
		allWords := readFile(inFile)
		CreateJsonFile(allWords, outFile)
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

func readFile(fileName string) ([]string) {
	var allWords []string
	log.WithFields(log.Fields{
		"file": fileName,
	}).Infoln("Parsing file")
	file,err := os.Open(fileName)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"filename": fileName,
		}).Errorln("Error opening file")
		return allWords
	}
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			log.Infoln("Reached end of file")
			break
		}

		str = strings.TrimSpace(str)

		allWords = append(allWords, str)
	}
	log.WithFields(log.Fields{
		"file": fileName,
		"rows": len(allWords),
	}).Infoln("Completed parsing file")

	return allWords
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&inFile, "input", "f", "dataset.txt", "name of the input file to read from")
	generateCmd.Flags().StringVarP(&outFile, "outfile", "o", "dataset.json", "name of the file to write to")
}
