package main

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestCanDetectRSTFormattingWhenPresent(t *testing.T) {
	retVal, _ := containsRSTFormatting(helperLoadBytes(t, "containsRSTFormatting.adoc"))
	if retVal != "Contains reStructuredText formatting" {
		t.Error("Did not properly detect rST formatting")
	}
}

func TestDoesNotDetectRSTFormattingWhenNotPresent(t *testing.T) {
	retVal, _ := containsRSTFormatting(helperLoadBytes(t, "containsNoRSTFormatting.adoc"))
	if retVal != "No reStructuredText formatting detected" {
		t.Error("Incorrectly detected rST formatting")
	}
}

func TestDoesNotCheckEmptyContent(t *testing.T) {
	retVal, err := containsRSTFormatting(helperLoadBytes(t, "empty-file.adoc"))
	if err == nil || retVal != "" {
		t.Error("Error should have been returned when attempting to process empty data")
	}
	if retVal != "" {
		t.Error("Return value should have been an empty string returned when attempting to process empty data")
	}
}

func helperLoadBytes(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}
