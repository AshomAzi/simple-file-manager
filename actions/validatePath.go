package actions

import "os"

func ValidatePath(path string) (bool, error) {

	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return true, nil
}
