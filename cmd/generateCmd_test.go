package cmd

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"github.com/harshc/indexconstructor/globals"
	log "github.com/Sirupsen/logrus"
)

var testWords = []string {
	"boinks",
	"boot",
	"booted",
	"booting",
	"boots",
	"bounce",
	"bounced",
	"bounces",
	"bouncing",
	"boustrophedon",
}

// BenchmarkCreateJsonFile ...
func BenchmarkCreateJsonFile(b *testing.B) {
	log.Infoln("Benchmark CreateJSON File")
	for n := 0; n < b.N; n++ {
		fileName := "benchmark.json"
		CreateJsonFile(testWords, fileName)
		os.Remove(fileName)
	}
}

// TestCreateJsonFile ...
func TestCreateJsonFile(t *testing.T) {
	log.Infoln("Test CreateJSONFile")
	fileName := "testing.json"

	log.Infoln("Create the file with testWords")
	CreateJsonFile(testWords, fileName)

	buff, err := ioutil.ReadFile(fileName)
	if err != nil {
		assert.Fail(t, "Error reading generated file ")
	}

	log.Infoln("Read the generatedFile and verify it is valid JSON")
	var scoredNames []globals.ScoredName
	if err := json.Unmarshal(buff, &scoredNames); err != nil {
		assert.Fail(t, "Incorrectly output json ")
	}

	assert.True(t, err==nil, "Valid JSON generated")

	expectedCount := len(testWords)
	actualCount := len(scoredNames)

	assert.Equal(t, expectedCount, actualCount, "JSON file correctly generated")

	os.Remove(fileName)
}