package cmd

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"github.com/harshc/indexconstructor/globals"
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
	for n := 0; n < b.N; n++ {
		fileName := "benchmark.json"
		CreateJsonFile(testWords, fileName)
		os.Remove(fileName)
	}
}

// TestCreateJsonFile ...
func TestCreateJsonFile(t *testing.T) {
	fileName := "testing.json"
	CreateJsonFile(testWords, fileName)

	buff, err := ioutil.ReadFile(fileName)
	if err != nil {
		assert.Fail(t, "Error reading generated file ")
	}

	var scoredNames []globals.ScoredName
	if err := json.Unmarshal(buff, &scoredNames); err != nil {
		assert.Fail(t, "Incorrectly output json ")
	}

	expectedCount := len(testWords)
	actualCount := len(scoredNames)

	assert.Equal(t, expectedCount, actualCount, "JSON file correctly generated")

	os.Remove(fileName)
}