package git

import (
	"os"
	"os/exec"
)

func Commit(message string) (string, error) {
	file, err := os.CreateTemp("", "GIT_HOOK")
	if err != nil {
		return "", err
	}
	defer os.Remove(file.Name())

	if _, err := file.Write([]byte(message)); err != nil {
		return "", err
	}

	cmd := exec.Command("git", "commit", "-F")
	cmd.Args = append(cmd.Args, file.Name())

	result, error := cmd.CombinedOutput()

	return string(result), error
}
