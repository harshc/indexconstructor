package cmd

import (
	"testing"
	"io"
	"os"
	"github.com/harshc/indexconstructor/globals"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	log "github.com/Sirupsen/logrus"
)

// TestSerializeJson ...
func TestSerializeJson(t *testing.T) {
	fileName := "testing.json"
	CreateJsonFile(testWords, fileName)
	pr, pw := io.Pipe()

	log.Infoln("Test SerializeJSON")
	SerializeJson(fileName, "", pw, pr, mockHttpCaller)
	log.Infoln("Serialized data sent over the pipe")
	expectedCount := len(testWords)
	go func() {
		log.Infoln("Validate that we received the data")
		actualCount := decodeJsonStream(pr)
		assert.Equal(t, expectedCount, actualCount , "JSON file correctly serialized and sent over the wire")
	}()

	os.Remove(fileName)
}

// func BenchmarkSerializeJson ...
func BenchmarkSerializeJson(b *testing.B) {
	fileName := "benchmark.json"
	CreateJsonFile(testWords, fileName)
	pr, pw := io.Pipe()

	for i := 0; i < b.N; i++ {
		SerializeJson(fileName, "", pw, pr, mockHttpCaller)
	}

	os.Remove(fileName)
}

func decodeJsonStream(pr *io.PipeReader) (int) {
	var scoredNames []globals.ScoredName
	log.Infoln("Lets decode in a go routine")
	decoder := json.NewDecoder(pr)
	err := decoder.Decode(&scoredNames)

	log.WithFields(log.Fields{
		"names": scoredNames,
	}).Infoln("Decoded names from the Pipereader")

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorln("Error decoding json from the reader")
		return 0
	}

	actualCount := len(scoredNames)

	return actualCount

}
func mockHttpCaller(url string, reader *io.PipeReader) {
// do nothing
}