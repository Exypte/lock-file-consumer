// Package yamlprint parses YAML files with viper and prints them back as YAML.
package yamlprint

import (
	"fmt"
	"io"

	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
)

// DefaultIndent is the indentation used when Options.Indent is zero.
const DefaultIndent = 2

// Options configures Print.
type Options struct {
	// Indent is the number of spaces used for indentation. Zero means
	// DefaultIndent; negative values are rejected.
	Indent int
}

// Print parses the YAML file at path and writes it to w as YAML.
//
// Parsing is done by viper, so map keys are lowercased and printed in
// alphabetical order, and comments are dropped.
func Print(w io.Writer, path string, opts Options) error {
	indent := opts.Indent
	switch {
	case indent == 0:
		indent = DefaultIndent
	case indent < 0:
		return fmt.Errorf("invalid indent %d: must not be negative", indent)
	}

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("parse %s: %w", path, err)
	}

	enc := yaml.NewEncoder(w)
	enc.SetIndent(indent)
	if err := enc.Encode(v.AllSettings()); err != nil {
		return fmt.Errorf("encode %s: %w", path, err)
	}
	return enc.Close()
}
