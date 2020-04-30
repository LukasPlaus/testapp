package main

import (
	"regexp"
	"testing"
)

// TestRevisionString - Tests if revision number matches agreed format
func TestRevisionString(t *testing.T) {
	appVersion := getVersion()

	expectedStringFormat := "v\\d{9}_\\d{6}"
	regexMatch := regexp.MustCompile(expectedStringFormat)

	matchFound := regexMatch.MatchString(appVersion)

	if matchFound == false {
		t.Errorf("Revision number format is incorrect, string: %v, does not match pattern: %v.", appVersion, expectedStringFormat)
	}

	// fmt.Printf("\nRevision string: %v :%v\n", appVersion, regexMatch.MatchString(appVersion))
}
