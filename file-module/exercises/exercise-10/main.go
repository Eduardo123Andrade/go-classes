package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dirPath := "../../files"

	filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}

		info, err := d.Info()

		if err != nil {
			panic(err)
		}

		now := time.Now().Add(-24 * time.Hour)

		if info.ModTime().Before(now) {
			fmt.Println("File modified more than 24 hours ago:", info.Name())

			os.Remove(path)

			return nil
		}
		return nil
	})
}
