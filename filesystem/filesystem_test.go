package filesystem

import (
	"os"
	"testing"
	"testing/fstest"
)

func TestCountFiles(t *testing.T) {

	tests := []struct {
		label string
		fs    fstest.MapFS
		count int
		err   error
	}{
		{
			"no files",
			fstest.MapFS{},
			0,
			nil,
		},
		{
			"one file",
			fstest.MapFS{"foo": {}},
			1,
			nil,
		},
		{
			"one file, one dir",
			fstest.MapFS{"foo": {}, "bar": {Mode: os.ModeDir}},
			1,
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.label, func(t *testing.T) {
			count, err := CountFiles(tc.fs)
			if tc.err != err {
				t.Errorf("incorrect err: %s expected %s", err, tc.err)
			}

			if count != tc.count {
				t.Errorf("incorrect count: %d expected %d", count, tc.count)
			}
		})
	}
}
