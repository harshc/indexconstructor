package cmd

import (
	"github.com/spf13/cobra"
	log "github.com/Sirupsen/logrus"
	"encoding/json"
	"os"
	"github.com/harshc/indexconstructor/globals"
	"io"
)

var infile string

type HttpPOST func(url string, pipeReader *io.PipeReader)

var processCmd = &cobra.Command {
	Use: "process",
	Short: "Process a json file and generate a serialized data structure",
	Long: "Process the json file and generate the serialized stream to a file",
	Run: func(cmd *cobra.Command, args []string) {
		pr, pw := io.Pipe()
		go SerializeJson(infile, "", pw, pr, callHTTPEndpoint)
	},
}

// SerializeJson ...
func SerializeJson(fileName string, url string, pipeWriter *io.PipeWriter, pipeReader *io.PipeReader, httpCaller HttpPOST) {
	log.Debugln("Starting to SerializeJSON")

	file, err := os.Open(fileName)
	log.WithFields(log.Fields{
		"file": fileName,
	}).Infoln("Opened filestream")

	defer file.Close()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"filename": fileName,
		}).Errorln("Error creating a file reader ")
	}

	// Create a json decoder
	decoder := json.NewDecoder(file)
	encoder := json.NewEncoder(pipeWriter)
	decodeAndStreamJson(decoder, encoder, pipeWriter)
	// we can concurrently post to the HTTP endpoint
	httpCaller(url, pipeReader)
}

func decodeAndStreamJson(decoder *json.Decoder, encoder *json.Encoder, pipeWriter *io.PipeWriter) {
	// we have to verify that the json is valid, before passing this along
	var names []globals.ScoredName
	err := decoder.Decode(&names)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorln("Error parsing token")
	} else {
		// turn around and stream this to the POST
		log.Infoln("Begin streaming")
		var name globals.ScoredName

		go func() {
			defer pipeWriter.Close()

			err := encoder.Encode(&names)
			log.WithFields(log.Fields{
				"names": names,
			}).Debugln("Encoded names")

			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
					"name":  name.Name,
					"score": name.Score,
				}).Errorln("Error encoding value to json stream")
			}
		}()
		log.Infoln("Done sending the encoded stream over the wire")
	}
}

func callHTTPEndpoint(canonicalurl string, pipeReader *io.PipeReader) {
 // do nothing here
}

func init() {
	rootCmd.AddCommand(processCmd)
	processCmd.Flags().StringVarP(&infile, "infile", "i", "mockdataset.json", "Input file with the dataset")
}
