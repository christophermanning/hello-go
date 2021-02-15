package filesystem

import (
	"io/fs"
)

func CountFiles(dfs fs.FS) (int, error) {
	num := 0

	fn := func(path string, d fs.DirEntry, err error) error {
		// only count files, but allow walking into the top level directory for counting
		if path != "." && d.IsDir() {
			return fs.SkipDir
		} else if !d.IsDir() {
			num++
		}

		return nil
	}

	fs.WalkDir(dfs, ".", fn)

	return num, nil
}
