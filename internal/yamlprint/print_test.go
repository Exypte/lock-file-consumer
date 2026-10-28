package yamlprint

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// writeFile writes content to a new file in a temporary directory and
// returns its path.
func writeFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "input.yaml")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("os.WriteFile(%q) failed: %v", path, err)
	}
	return path
}

func TestPrint(t *testing.T) {
	tests := []struct {
		name  string
		input string
		opts  Options
		want  string
	}{
		{
			name:  "nested map and list",
			input: "name: demo\nsteps:\n  - checkout\n  - build\nenv:\n  go: \"1.27\"\n",
			want:  "env:\n  go: \"1.27\"\nname: demo\nsteps:\n  - checkout\n  - build\n",
		},
		{
			name:  "keys are lowercased and sorted",
			input: "Zeta: 1\nAlpha: 2\n",
			want:  "alpha: 2\nzeta: 1\n",
		},
		{
			name:  "custom indent",
			input: "a:\n  b: c\n",
			opts:  Options{Indent: 4},
			want:  "a:\n    b: c\n",
		},
		{
			name:  "empty file",
			input: "",
			want:  "{}\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := writeFile(t, tc.input)
			var got strings.Builder
			if err := Print(&got, path, tc.opts); err != nil {
				t.Fatalf("Print(%q, %+v) failed: %v", tc.input, tc.opts, err)
			}
			if diff := cmp.Diff(tc.want, got.String()); diff != "" {
				t.Errorf("Print(%q, %+v) output mismatch (-want +got):\n%s", tc.input, tc.opts, diff)
			}
		})
	}
}

func TestPrintErrors(t *testing.T) {
	tests := []struct {
		name string
		path string
		opts Options
	}{
		{
			name: "missing file",
			path: filepath.Join(t.TempDir(), "missing.yaml"),
		},
		{
			name: "invalid yaml",
			path: writeFile(t, "a: [b\n"),
		},
		{
			name: "negative indent",
			path: writeFile(t, "a: b\n"),
			opts: Options{Indent: -1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var out strings.Builder
			if err := Print(&out, tc.path, tc.opts); err == nil {
				t.Errorf("Print(%q, %+v) = nil, want error", tc.path, tc.opts)
			}
		})
	}
}
