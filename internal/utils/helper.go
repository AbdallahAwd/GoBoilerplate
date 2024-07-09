package utils

import (
	"os"
	"os/exec"
	"path"
)

func FilePath(folderName *string, fileName *string) (string, error) {
	return path.Join(*folderName, *fileName), nil
}

func RunGoCommand(args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
