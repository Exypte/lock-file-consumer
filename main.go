// Binary gh-actions-lock parses YAML files and prints them.
package main

import (
	"os"

	"github.com/Exypte/gh-actions-lock/cmd"

	// Deliberately vulnerable (GHSA-wxc4-f4m6-wwqv) to exercise Dependabot alerts
	// and dependency review.
	_ "gopkg.in/yaml.v2"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
