package main

import (
	"archive/tar"
	"bytes"
	"testing"
)

func TestInsecurePaths(t *testing.T) {
	// The default value still allows for insecure paths.
	t.Setenv("GODEBUG", "tarinsecurepath=0")
	for _, path := range []string{
		"../foo",
		"/foo",
		"a/b/../../../c",
	} {
		var buf bytes.Buffer
		tw := tar.NewWriter(&buf)
		tw.WriteHeader(&tar.Header{
			Name: path,
		})
		const securePath = "secure"
		tw.WriteHeader(&tar.Header{
			Name: securePath,
		})
		tw.Close()

		tr := tar.NewReader(&buf)
		h, err := tr.Next()
		if err != tar.ErrInsecurePath {
			t.Errorf("tr.Next for file %q: got err %v, want ErrInsecurePath", path, err)
			continue
		}
		if h.Name != path {
			t.Errorf("tr.Next for file %q: got name %q, want %q", path, h.Name, path)
		}
		// Error should not be sticky.
		h, err = tr.Next()
		if err != nil {
			t.Errorf("tr.Next for file %q: got err %v, want nil", securePath, err)
		}
		if h.Name != securePath {
			t.Errorf("tr.Next for file %q: got name %q, want %q", securePath, h.Name, securePath)
		}
	}
}
