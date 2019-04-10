package init

import (
	"os"
)

func Exec() error {

	if err := os.MkdirAll("controllers", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("entities/models", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("entities/repositories", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("usecases", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("usecases/inputs/", 0777); err != nil {
		return err
	}

	if err := os.MkdirAll("usecases/outputs", 0777); err != nil {
		return err
	}

	return nil
}
