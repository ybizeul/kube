package cmd

import (
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Copy file to another folder
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
