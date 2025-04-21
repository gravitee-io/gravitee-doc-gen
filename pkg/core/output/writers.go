package output

import (
	"fmt"
	"os"
	"path/filepath"
)

type Console struct {
	To string
}
type File struct {
	To string
}

func (f File) Write(generated []byte) (int, error) {

	stat, err := os.Stat(f.To)

	if err != nil && os.IsNotExist(err) {
		dir := filepath.Dir(f.To)
		err := os.MkdirAll(dir, 0744)
		if err != nil {
			return 0, err
		}
		newFile, err := os.Create(f.To)
		if err != nil {
			return 0, err
		}

		stat, err = newFile.Stat()
		if err != nil {
			return 0, err
		}

	} else if err != nil {
		return 0, err
	}

	err = os.WriteFile(f.To, generated, stat.Mode().Perm())
	if err != nil {
		return 0, err
	}

	return len(generated), nil

}

func (c Console) Write(generated []byte) (int, error) {
	fmt.Println("--- Dry Run (" + c.To + ") ---")
	fmt.Println(string(generated))
	fmt.Println("-----------------------")
	fmt.Println()
	return len(generated), nil
}
