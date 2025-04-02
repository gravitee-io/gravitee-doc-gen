package output

import (
	"fmt"
	"os"
)

type Console struct{}
type File struct{}

func (f File) Write(generated []byte) (int, error) {
	stat, err := os.Stat(readmeFileName)
	if err != nil {
		return 0, err
	}

	err = os.WriteFile(readmeFileName, generated, stat.Mode().Perm())
	if err != nil {
		return 0, err
	}

	return len(generated), nil

}

func (c Console) Write(generated []byte) (int, error) {
	fmt.Println("--- Dry Run ---")
	fmt.Println(string(generated))
	fmt.Println("---------------")
	return len(generated), nil
}
