package git

import (
	"log"
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

func Hook(message string, path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("run git commit failed, err=%v\n", err)
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		log.Printf("run git commit failed, err=%v\n", err)
	}
}
